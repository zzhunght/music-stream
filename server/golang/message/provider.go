package message

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProvider struct {
	amqpConn *amqp.Connection
	ch       *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQProvider, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMQProvider{
		amqpConn: conn,
		ch:       ch,
	}, nil
}

func (r *RabbitMQProvider) CloseRabbitMQ() error {
	defer r.amqpConn.Close()
	return r.ch.Close()
}

func (r *RabbitMQProvider) Send() error {
	err := r.ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}
	return nil
}
