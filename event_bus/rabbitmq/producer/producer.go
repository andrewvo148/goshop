package producer

import (
	"context"
	"encoding/json"
	"event-bus/event_bus"

	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConnection(url string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, errors.Wrap(err, "amqp.Dial")
	}

	return conn, nil
}

func NewProducer(conn *amqp.Connection, opts ...Option) (*producer, error) {
	p := &producer{
		conn: conn,
	}
	for _, opt := range opts {
		opt(p)
	}

	return p, nil
}

func (p *producer) Publish(ctx context.Context, event event_bus.IntegrationEvent) error {
	ch, err := p.conn.Channel()
	if err != nil {
		return errors.Wrap(err, "CreateChannel")
	}

	err = ch.ExchangeDeclare(
		p.exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "ExchangeDeclare")
	}

	body, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "Marshalling")
	}

	err = ch.PublishWithContext(ctx,
		p.exchangeName, // exchange name
		p.bindingKey,   // routing key
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		})

	if err != nil {
		return errors.Wrap(err, "publish")
	}

	return nil
}
