package main

import (
	"fmt"
	"net/http"
)

// "github.com/arturbaccarin/credit-card-messaging/register-service/internal/database"

func main() {
	println("Starting register service")
	// db := database.StartDB()

	resp, err := http.Get("http://localhost:3000/receive")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
