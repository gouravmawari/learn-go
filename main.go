package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"time"
)
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var jwtSecret = []byte("my-secret-key");

func generatetoken(username string)(string,error){
claim := jwt.MapClain{
	"username":username,
	"exp":time.Now().Add(time.Hour * 24).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim);
	tokenString, err := token.SignedString(jwtSecret)
	return token,err
}

func NewPost(w http.ResponseWriter,r *http.Request){
	if r.Method != "POST" {
		http.Error(w,"this method is POST",http.StatusMethodNotAllowed)
		return
	}
	var user User;
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w,"invlide JSON",http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func main(){
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w ,"this is response");
	})
	http.HandleFunc("/user",NewPost);
	http.ListenAndServe(":4563",nil)

}