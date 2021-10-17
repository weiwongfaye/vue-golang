package endpoints

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getSecret() string {
	secret := os.Getenv("ACCESS_SECRET")
	if secret == "" {
		//That's surely a big secret this way...
		secret = "sdmalncnjsdsmf"
	}
	return secret
}

var users map[string][]byte = make(map[string][]byte)
var idxUsers int = 0

//getTokenUserPassword returns a jwt token for a user if the //password is ok
func getTokenUserPassword(w http.ResponseWriter, r *http.Request) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "cannot decode username/password struct", http.StatusBadRequest)
		return
	}
	//here I have a user!
	//Now check if exists
	passwordHash, found := users[u.Username]
	if !found {
		http.Error(w, "Cannot find the username", http.StatusNotFound)
	}
	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(u.Password))
	if err != nil {
		return
	}
	token, err := createToken(u.Username)
	if err != nil {
		http.Error(w, "Cannot create token", http.StatusInternalServerError)
		return
	}
	sendJSONResponse(w, struct {
		Token string `json:"token"`
	}{token})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Cannot decode request", http.StatusBadRequest)
		return
	}
	if _, found := users[u.Username]; found {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	//If I'm here-> add user and return a token
	value, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	users[u.Username] = value
	token, err := createToken(u.Username)
	if err != nil {
		http.Error(w, "Cannot create token", http.StatusInternalServerError)
		return
	}
	sendJSONResponse(w, struct {
		Token string `json:"token"`
	}{token})
}
func createToken(username string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	secret := getSecret()
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func getTokenByToken(w http.ResponseWriter, r *http.Request) {
	//Here I already have the token checked... Just get the username from Request context
	username, ok := context.Get(r, "username").(string)
	if !ok {
		http.Error(w, "Cannot check username", http.StatusInternalServerError)
		return
	}
	token, err := createToken(username)
	if err != nil {
		http.Error(w, "Cannot create token", http.StatusInternalServerError)
		return
	}
	sendJSONResponse(w, struct{ Token string }{token})
}
