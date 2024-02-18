package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arturbaccarin/credit-card-messaging/register-service/internal/database"
	"github.com/arturbaccarin/credit-card-messaging/register-service/internal/database/repository"
	"github.com/arturbaccarin/credit-card-messaging/register-service/internal/model"
	"gorm.io/gorm"
)

func main() {
	println("Starting register service")

	db := database.StartDB()
	processAuditMessages(db)
	processPaymentMessages(db)
}

func processAuditMessages(db *gorm.DB) {
	messages, err := receiveMessages("a")
	if err != nil {
		println(err)
	}

	err = registerAuditOrders(db, messages)
	if err != nil {
		println(err)
	}
}

func processPaymentMessages(db *gorm.DB) {
	messages, err := receiveMessages("r")
	if err != nil {
		println(err)
	}

	err = registerPaymentOrder(db, messages)
	if err != nil {
		println(err)
	}
}

func receiveMessages(key string) ([]string, error) {
	url := "http://localhost:3000/receive?k=" + key

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error receiving message: %q", err)
	}
	defer resp.Body.Close()

	var messages model.Message
	err = json.NewDecoder(resp.Body).Decode(&messages)
	if err != nil {
		return nil, fmt.Errorf("error decoding message: %q", err)
	}

	return messages.Messages, nil
}

func registerAuditOrders(db *gorm.DB, messages []string) error {
	repo := repository.NewAuditOrder(db)

	for _, message := range messages {
		order := model.PaymentOrderMessage{}
		err := json.Unmarshal([]byte(message), &order)
		if err != nil {
			fmt.Printf("error audit unmarshal message: %q", err)
			continue
		}

		err = repo.Create(order.ID, order.Date)
		if err != nil {
			fmt.Printf("error audit creating register: %q", err)
			continue
		}
	}

	return nil
}

func registerPaymentOrder(db *gorm.DB, messages []string) error {
	repo := repository.NewPaymentOrder(db)

	for _, message := range messages {
		order := model.PaymentOrderMessage{}
		err := json.Unmarshal([]byte(message), &order)
		if err != nil {
			fmt.Printf("error payment unmarshal message: %q", err)
			continue
		}

		err = repo.Create(order.ID, order.Value, order.Date)
		if err != nil {
			fmt.Printf("error payment creating register: %q", err)
			continue
		}
	}

	return nil
}
