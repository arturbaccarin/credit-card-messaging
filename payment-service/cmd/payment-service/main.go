package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arturbaccarin/credit-card-messaging/payment-service/internal/generator"
	"github.com/arturbaccarin/credit-card-messaging/payment-service/internal/model"
)

func main() {
	paymentOrderGenerator := generator.NewPaymentOrderGenerator(3)

	paymentOrders := paymentOrderGenerator.GenerateCopies()

	for _, paymentOrder := range paymentOrders {
		message, err := json.Marshal(paymentOrder)
		if err != nil {
			fmt.Println("Error marshalling struct to JSON:", err)
			continue
		}

		body := model.RabbitMQPayload{
			Message: string(message),
		}

		payload, err := json.Marshal(body)
		if err != nil {
			fmt.Println("Error marshalling struct to JSON:", err)
			continue
		}

		_, err = http.Post("http://localhost:3000/send", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error sending message:", err)
		}

		fmt.Println("Message sent successfully")
	}
}
