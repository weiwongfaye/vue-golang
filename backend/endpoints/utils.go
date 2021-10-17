package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func AddRouterEndpoints(r *mux.Router) *mux.Router {
	r.HandleFunc("/api/posts", getPosts).Methods("GET")
	r.HandleFunc("/api/posts", checkTokenHandler(addPost)).Methods("POST")
	r.HandleFunc("/api/posts/{POST_ID}", checkTokenHandler(deletePost)).Methods("Delete")
	r.HandleFunc("/api/posts/{POST_ID}/comments", checkTokenHandler(addComment)).Methods("POST")
	r.HandleFunc("/api/auth/login", getTokenUserPassword).Methods("POST")
	r.HandleFunc("/api/auth/create-user", createUser).Methods("POST")
	r.HandleFunc("/api/auth/refresh-token", checkTokenHandler(getTokenByToken)).Methods("GET")
	r.HandleFunc("/api/users/{USERNAME}", checkTokenHandler(getUser)).Methods("GET")
	return r
}

func checkTokenHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		bearerToken := strings.Split(header, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Cannot read token", http.StatusBadRequest)
			return
		}
		if bearerToken[0] != "Bearer" {
			http.Error(w, "Error in authorization token. it needs to be in form of 'Bearer <token>'", http.StatusBadRequest)
			return
		}
		token, ok := checkToken(bearerToken[1])
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			username, ok := claims["username"].(string)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			//check if username actually exists
			if _, ok := users[username]; !ok {
				http.Error(w, "Unauthorized, user not exists", http.StatusUnauthorized)
				return
			}
			//Set the username in the request, so I will use it in check after!
			context.Set(r, "username", username)
		}
		next(w, r)
	}
}

func checkToken(tokenString string) (*jwt.Token, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getSecret()), nil
	})
	if err != nil {
		return nil, false
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, false
	}
	return token, true
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write the response body: %v", err)
		return
	}

}
