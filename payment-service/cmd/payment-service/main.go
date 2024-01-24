package main

import (
	"encoding/json"
	"fmt"

	"github.com/arturbaccarin/credit-card-messaging/payment-service/internal/generator"
)

func main() {
	paymentOrderGenerator := generator.NewPaymentOrderGenerator(3)

	paymentOrders := paymentOrderGenerator.GenerateCopies()

	for _, paymentOrder := range paymentOrders {
		jsonData, err := json.Marshal(paymentOrder)
		if err != nil {
			fmt.Println("Error marshalling struct to JSON:", err)
			continue
		}

		requestBody := struct {
			Message string
		}{
			string(jsonData),
		}

	}

}
