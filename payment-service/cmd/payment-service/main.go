package main

import (
	"fmt"

	"github.com/arturbaccarin/credit-card-messaging/payment-service/internal/generator"
)

func main() {
	paymentOrderGenerator := generator.NewPaymentOrderGenerator(3)

	paymentOrders := paymentOrderGenerator.GenerateCopies()

	for _, t := range paymentOrders {
		fmt.Println(t)
	}
}
