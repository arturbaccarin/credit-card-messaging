package rabbitmq

import (
	"log"
	"sync"

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

	var wg sync.WaitGroup

	go func() {
		for {
			wg.Add(1)
			msg, ok := <-msgs
			if !ok {
				wg.Done()
				break
			}

			log.Printf(" > Received message: %s\n", msg.Body)
			wg.Done()
		}
	}()

	wg.Wait()

	return messages, nil
}
