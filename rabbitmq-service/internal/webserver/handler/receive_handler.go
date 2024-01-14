package handler

import (
	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ReceiveHandler struct {
	ch *amqp.Channel
}

func NewReceiveHandler(ch *amqp.Channel) *ReceiveHandler {
	return &ReceiveHandler{
		ch: ch,
	}
}

func (r ReceiveHandler) ReceiveMessages(c *fiber.Ctx) error {
	messages, err := rabbitmq.Consume(r.ch)
	if err != nil {
		return err
	}

	_, err = c.Write(messages)
	return err
}
