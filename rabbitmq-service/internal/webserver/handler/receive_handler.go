package handler

import (
	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ReceiveHandler struct {
	conn *amqp.Connection
}

func NewReceiveHandler(conn *amqp.Connection) *ReceiveHandler {
	return &ReceiveHandler{
		conn: conn,
	}
}

func (r ReceiveHandler) ReceiveMessages(c *fiber.Ctx) error {
	ch, err := rabbitmq.NewChannel(r.conn)
	if err != nil {
		return err
	}
	defer ch.Close()

	messages, err := rabbitmq.Consume(ch)
	if err != nil {
		return err
	}

	_, err = c.Write(messages)
	return err
}
