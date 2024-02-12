package handler

import (
	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/internal/dto"
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
	queue := c.Query("q")

	messages, err := rabbitmq.Consume(r.ch, queue)
	if err != nil {
		return err
	}

	var res dto.MessagesResponse
	for _, message := range messages {
		res.Messages = append(res.Messages, string(message))
	}

	return c.JSON(res)
}
