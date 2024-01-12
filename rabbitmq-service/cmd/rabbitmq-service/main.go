package main

import (
	"github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/internal/rabbit"
)

func main() {
	conn := rabbit.NewConnection()
	ch := rabbit.NewChannel(conn)

	rabbit.Publish(ch)
	rabbit.Consume(ch)
}
