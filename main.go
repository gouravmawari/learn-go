package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var jwtSecret = []byte("my-secret-key")

func GenerateToken(username string) (string, error) {

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		authHead := r.Header.Get("Authorization")

		if authHead == "" {
			http.Error(w, "missing header", http.StatusUnauthorized)
			return
		}

		tokenString := authHead[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func NewPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "this method is POST", http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func main() {

	http.HandleFunc("/", JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is response")
	}))

	http.HandleFunc("/user", NewPost)

	http.ListenAndServe(":4563", nil)
}