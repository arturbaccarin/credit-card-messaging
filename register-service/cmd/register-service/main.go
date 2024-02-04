package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arturbaccarin/credit-card-messaging/register-service/internal/database"
	"github.com/arturbaccarin/credit-card-messaging/register-service/internal/database/repository"
	"github.com/arturbaccarin/credit-card-messaging/register-service/internal/model"
)

// "github.com/arturbaccarin/credit-card-messaging/register-service/internal/database"

func main() {
	println("Starting register service")
	db := database.StartDB()

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

	paymentOrderRepository := repository.NewPaymentOrder(db)

	for _, message := range messages.Messages {
		paymentOrderMessage := model.PaymentOrderMessage{}
		err = json.Unmarshal([]byte(message), &paymentOrderMessage)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = paymentOrderRepository.Create(paymentOrderMessage.ID, paymentOrderMessage.Value, paymentOrderMessage.Date)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
