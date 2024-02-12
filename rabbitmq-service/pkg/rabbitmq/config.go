package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	return conn
}

func NewChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	return ch
}

func NewDirectExchange(ch *amqp.Channel, name string) {
	err := ch.ExchangeDeclare(name, "direct", true, false, false, false, nil)
	if err != nil {
		log.Panicf("%s: %s", "Failed to create an exchange", err)
	}
}

func NewQueue(ch *amqp.Channel, name string) (amqp.Queue, error) {
	return ch.QueueDeclare(name, true, false, false, false, nil)
}

func NewBind(ch *amqp.Channel, queue *amqp.Queue, key, exchange string) error {
	return ch.QueueBind(queue.Name, key, exchange, false, nil)
}
