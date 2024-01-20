package main

import (
	"fmt"

	"github.com/arturbaccarin/credit-card-messaging/register-service/internal/database"
)

func main() {
	println("Starting register service")
	db := database.StartDB()
	fmt.Println(db)
}
