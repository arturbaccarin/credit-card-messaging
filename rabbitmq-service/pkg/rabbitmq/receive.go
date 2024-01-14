package rabbitmq

import (
	"log"
	"sync"
	"time"

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
	// done := make(chan bool)
	wg.Add(1)

	go func() {
		log.Println("Here I am in time 1")
		time.Sleep(2 * time.Second)
		wg.Done()
		log.Println("Here I am in time 2")
	}()

	go func() {
		for msg := range msgs {
			log.Printf(" > Received message: %s\n", msg.Body)
			wg.Done()
			break
		}
		log.Println("out1")
	}()

	wg.Wait()
	log.Println("out2")
	ch.Close()

	return messages, nil
}
