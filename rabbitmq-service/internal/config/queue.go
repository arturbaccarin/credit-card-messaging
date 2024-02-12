package config

import (
	"log"

	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewQueue(ch *amqp.Channel, name, exchange string, keys []string) {
	queue, err := rabbitmq.NewQueue(ch, name)
	if err != nil {
		log.Panicf("%s: %s", "Failed to create a queue", err)
	}

	for _, key := range keys {
		err = rabbitmq.NewBind(ch, &queue, key, exchange)
		if err != nil {
			log.Panicf("%s: %s", "Failed to bind a queue", err)
		}
	}
}
