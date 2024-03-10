package main

import "github.com/rafaelpapastamatiou/goexpert/11-events/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	rabbitmq.Publish(
		"amq.direct",
		"",
		ch,
		"Hello World!",
	)
}
