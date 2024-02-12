package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume(ch *amqp.Channel, queue string) ([]string, error) {
	var messages []string

	for {
		msg, ok, _ := ch.Get(queue, true)
		if !ok {
			break
		}
		messages = append(messages, string(msg.Body))
	}

	return messages, nil
}
