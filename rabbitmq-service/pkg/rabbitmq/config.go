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

func NewChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	return conn.Channel()
}

func NewDirectExchange(ch *amqp.Channel, name string) error {
	return ch.ExchangeDeclare(name, "direct", true, false, false, false, nil)
}

func NewQueue(ch *amqp.Channel, name string) (amqp.Queue, error) {
	return ch.QueueDeclare(name, true, false, false, false, nil)
}

func NewBind(ch *amqp.Channel, queue *amqp.Queue, key, exchange string) error {
	return ch.QueueBind(queue.Name, key, exchange, false, nil)
}
