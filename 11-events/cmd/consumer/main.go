package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rafaelpapastamatiou/goexpert/11-events/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume("fila", "consumer01", ch, msgs)

	for msg := range msgs {
		fmt.Println("Received message:", string(msg.Body))
		msg.Ack(false)
	}
}
