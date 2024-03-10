package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func OpenChannel() (*amqp.Channel, error) {
	conn, err := OpenConnection()
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func Consume(
	queue string,
	consumerName string,
	channel *amqp.Channel,
	outChan chan<- amqp.Delivery,
) error {
	msgs, err := channel.Consume(
		queue,
		consumerName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		outChan <- msg
	}

	return nil
}

func Publish(
	exchange string,
	routingKey string,
	channel *amqp.Channel,
	body string,
) error {
	err := channel.PublishWithContext(
		context.TODO(),
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
