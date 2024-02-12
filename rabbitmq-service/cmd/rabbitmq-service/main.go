package main

import (
	"log"

	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/internal/config"
	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/internal/webserver/handler"
	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	ampq "github.com/rabbitmq/amqp091-go"
)

var conn *ampq.Connection
var ch *ampq.Channel

func init() {
	conn = rabbitmq.NewConnection()
	ch = rabbitmq.NewChannel(conn)
	rabbitmq.NewDirectExchange(ch, "exchange")
	config.NewQueue(ch, "register", "exchange", []string{"r", "a"})
	config.NewQueue(ch, "audit", "exchange", []string{"a"})
}

func main() {
	app := fiber.New()
	app.Use(
		logger.New(),
	)

	sendHandler := handler.NewSendHandler(ch)
	receiveHandler := handler.NewReceiveHandler(ch)

	app.Post("/send", sendHandler.SendMessage)
	app.Get("/receive", receiveHandler.ReceiveMessages)

	log.Fatal(app.Listen(":3000"))
}
