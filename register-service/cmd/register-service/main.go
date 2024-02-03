package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arturbaccarin/credit-card-messaging/register-service/internal/model"
)

// "github.com/arturbaccarin/credit-card-messaging/register-service/internal/database"

func main() {
	println("Starting register service")
	// db := database.StartDB()

	resp, err := http.Get("http://localhost:3000/receive")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var messages model.Message
	err = json.NewDecoder(resp.Body).Decode(&messages)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(messages)
}
