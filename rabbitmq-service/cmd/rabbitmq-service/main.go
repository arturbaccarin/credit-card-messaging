package main

import "github.com/arturbaccarin/credit-card-messaging/rabbitmq-service/pkg/rabbitmq"

func main() {
	conn := rabbitmq.NewConnection()
	ch := rabbitmq.NewChannel(conn)

	rabbitmq.Publish(ch)
	rabbitmq.Consume(ch)
}
