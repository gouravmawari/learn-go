package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func UserData() string {
	time.Sleep(3 * time.Second)
	return "userData"
}

func OrderData() string {
	time.Sleep(3 * time.Second)
	return "orderData"
}

func notific() string {
	time.Sleep(3 * time.Second)
	return "notification"
}

type RoutineResponse struct {
	User         string `json:"user"`
	Order        string `json:"order"`
	Notification string `json:"notification"`
}

func GoRoutineHandler(w http.ResponseWriter, r *http.Request) {
	userChan := make(chan string)
	orderChan := make(chan string)
	notiChan := make(chan string)

	go func() { userChan <- UserData() }()
	go func() { orderChan <- OrderData() }()
	go func() { notiChan <- notific() }()

	result := RoutineResponse{
		User:         <-userChan,
		Order:        <-orderChan,
		Notification: <-notiChan,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	fmt.Println("GoRoutineHandler done")
}

func sendEmail(email string ){
	fmt.Println("email sent",email);
}

var jobs = make(chan string)
func emailWorker(){
	for email := range jobs{
		sendEmail(email);
	}

	// for range jobs{
	// 	sendEmail(email);
	// }
}
func register(w http.ResponseWriter,r *http.Request){
	jobs <- "smawari1000@gmail.com";
	// jobs <- struct{}{} 
	w.Write([]byte("User created"))
}