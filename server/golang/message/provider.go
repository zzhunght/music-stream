package message

import (
	"context"
	"fmt"
	"music-app-backend/internal/app/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProvider struct {
	amqpConn *amqp.Connection
	ch       *amqp.Channel
	config   *utils.Config
}

func NewRabbitMQ(config *utils.Config) (*RabbitMQProvider, error) {
	conn, err := amqp.Dial(config.RabbitMQUrl)
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
		config:   config,
	}, nil
}

func (r *RabbitMQProvider) CloseRabbitMQ() error {
	defer r.amqpConn.Close()
	return r.ch.Close()
}
func (r *RabbitMQProvider) DeclareExchange() error {
	err := r.ch.ExchangeDeclare(
		r.config.ExchangeName, // name
		"direct",              // type
		true,                  // durable
		false,                 // auto-deleted
		false,                 // internal
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *RabbitMQProvider) Publishing(body []byte) error {
	err := r.ch.PublishWithContext(context.Background(),
		r.config.ExchangeName,   // exchange
		r.config.NotiRoutingKey, // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		fmt.Print("error when sen message to queue", err)
	}

	return err
}
