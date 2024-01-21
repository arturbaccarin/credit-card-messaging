package handler

import (
	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/internal/dto"
	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

type SendHandler struct {
	conn *amqp.Connection
}

func NewSendHandler(conn *amqp.Connection) *SendHandler {
	return &SendHandler{
		conn: conn,
	}
}

func (s SendHandler) SendMessage(c *fiber.Ctx) error {
	ch, err := rabbitmq.NewChannel(s.conn)
	if err != nil {
		return err
	}

	var message dto.MessagePayload

	err = c.BodyParser(&message)
	if err != nil {
		return err
	}

	rabbitmq.Publish(ch, message.Message)

	return nil
}
