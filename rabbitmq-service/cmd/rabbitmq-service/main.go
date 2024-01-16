package main

import (
	"log"

	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/internal/webserver/handler"
	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	conn := rabbitmq.NewConnection()

	app := fiber.New()
	app.Use(
		logger.New(),
	)

	sendHandler := handler.NewSendHandler(conn)
	receiveHandler := handler.NewReceiveHandler(conn)

	app.Post("/send", sendHandler.SendMessage)
	app.Get("/receive", receiveHandler.ReceiveMessages)

	// https://dev.to/koddr/working-with-rabbitmq-in-golang-by-examples-2dcn
	// rabbitmq.Publish(ch)
	// rabbitmq.Consume(ch)
	log.Fatal(app.Listen(":3000"))
}
