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
