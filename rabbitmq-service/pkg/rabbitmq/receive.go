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
	wg.Add(10)

	go func() {
		for {
			select {
			case <-msgs:
				log.Println("ok")
			default:
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()

	return messages, nil
}
