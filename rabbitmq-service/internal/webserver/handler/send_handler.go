package handler

import (
	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/internal/dto"
	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

type SendHandler struct {
	ch *amqp.Channel
}

func NewSendHandler(ch *amqp.Channel) *SendHandler {
	return &SendHandler{
		ch: ch,
	}
}

func (s SendHandler) SendMessage(c *fiber.Ctx) error {
	var message dto.MessagePayload

	err := c.BodyParser(&message)
	if err != nil {
		return err
	}

	rabbitmq.Publish(s.ch, message.Message)

	return nil
}
