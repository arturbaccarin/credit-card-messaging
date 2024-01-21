package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume(ch *amqp.Channel) ([]byte, error) {
	defer ch.Close()
	var messages []byte

	for {
		msg, ok, _ := ch.Get("hello", true)
		if !ok {
			break
		}
		messages = append(messages, msg.Body...)
	}

	return messages, nil
}
