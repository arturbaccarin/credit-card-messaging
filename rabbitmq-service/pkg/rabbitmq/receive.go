package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume(ch *amqp.Channel) ([]byte, error) {
	var messages []byte

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		return nil, err
	}

	go func() {
		for {
			msg, ok := <-msgs
			if !ok {
				break
			}
			log.Printf(" > Received message: %s\n", msg.Body)
		}
	}()

	ch.Close()

	return messages, nil
}
