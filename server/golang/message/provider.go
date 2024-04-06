package message

import (
	"context"

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
func (r *RabbitMQProvider) DeclareExchange() error {
	err := r.ch.ExchangeDeclare(
		"music_create", // name
		"fanout",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *RabbitMQProvider) Send(ctx context.Context) error {
	err := r.ch.PublishWithContext(ctx,
		"music_create", // exchange
		"",             // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello worlds"),
		})

	return err
}
