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

	sendindOrders(paymentOrders)
}

func sendindOrders(paymentOrders []model.PaymentOrder) {
	for _, paymentOrder := range paymentOrders {
		payload, err := parseMessage(paymentOrder)
		if err != nil {
			continue
		}

		sendMessage(paymentOrder.NeedAudit, payload)
	}
}

func parseMessage(paymentOrder model.PaymentOrder) ([]byte, error) {
	message, err := json.Marshal(paymentOrder)
	if err != nil {
		fmt.Println("Error marshalling struct to JSON:", err)
		return nil, err
	}

	body := model.RabbitMQPayload{
		Message: string(message),
	}

	payload, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshalling struct to JSON:", err)
		return nil, err
	}

	return payload, nil
}

func sendMessage(needAudit bool, payload []byte) {
	var key string

	if needAudit {
		key = "a"
	} else {
		key = "r"
	}

	url := fmt.Sprintf("http://localhost:3000/messages?k=%s", key)

	_, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
