package main

import (
	"github.com/arturbaccarin/credit-card-messaging/internal/rabbit"
)

func main() {
	conn := rabbit.NewConnection()
	ch := rabbit.NewChannel(conn)

	rabbit.Publish(ch)
}
