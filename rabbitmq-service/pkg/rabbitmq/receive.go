package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume(ch *amqp.Channel) ([]string, error) {
	defer ch.Close()
	var messages []string

	for {
		msg, ok, _ := ch.Get("hello", true)
		if !ok {
			break
		}
		messages = append(messages, string(msg.Body))
	}

	return messages, nil
}
