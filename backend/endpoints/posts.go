package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type comment struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Post     string    `json:"post"`
	Date     time.Time `json:"date"`
}

type post struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Post     string    `json:"post"`
	Date     time.Time `json:"date"`
	Comments []comment `json:"comments"`
}

var posts []post = make([]post, 0)
var index int = 1

func addPost(w http.ResponseWriter, r *http.Request) {
	log.Println("addPost called")
	var actualPost post
	err := json.NewDecoder(r.Body).Decode(&actualPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	actualPost.ID = index
	index++
	actualPost.Date = time.Now()
	if actualPost.Comments == nil {
		actualPost.Comments = make([]comment, 0)
	}
	posts = append(posts, actualPost)
	sendJSONResponse(w, actualPost)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	log.Println("deletePost called")
	vars := mux.Vars(r)
	idString, ok := vars["POST_ID"]
	if !ok {
		http.Error(w, "Cannot find ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Cannot convert the id value to string", http.StatusBadRequest)
		return
	}
	for i := 0; i < len(posts); i++ {
		if posts[i].ID == id {
			posts[i] = posts[len(posts)-1]
			posts = posts[:len(posts)-1]
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, "Cannot find the requested id", http.StatusNotFound)

}

func addComment(w http.ResponseWriter, r *http.Request) {
	log.Println("addComment called")
	vars := mux.Vars(r)
	idString, ok := vars["POST_ID"]
	if !ok {
		http.Error(w, "Cannot find PostID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Cannot convert the POST_ID value to string", http.StatusBadRequest)
		return
	}
	var actualComment comment
	err = json.NewDecoder(r.Body).Decode(&actualComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i := 0; i < len(posts); i++ {
		if posts[i].ID == id {
			var commMax int = 0
			var commentIDs []int = []int{0}

			for _, comment := range posts[i].Comments {
				commentIDs = append(commentIDs, comment.ID)
			}
			sort.Ints(commentIDs)
			commMax = commentIDs[len(commentIDs)-1]
			actualComment.ID = commMax + 1
			actualComment.Date = time.Now()
			posts[i].Comments = append(posts[i].Comments, actualComment)
			sendJSONResponse(w, posts[i])
			return
		}
	}
	http.Error(w, "Cannot find a post with the selected id", http.StatusNotFound)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("GET post called")
	sendJSONResponse(w, posts)
}
