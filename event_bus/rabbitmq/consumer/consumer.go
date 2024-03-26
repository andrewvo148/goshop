package consumer

import (
	"context"
	"encoding/json"
	"event-bus/event_bus"

	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConsumer(conn *amqp.Connection, opts ...Option) *consumer {
	c := &consumer{
		conn: conn,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *consumer) Subcribe(ctx context.Context, handler event_bus.IntegrationEventHandler) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return errors.Wrap(err, "c.conn.Channel")
	}

	err = ch.ExchangeDeclare(
		c.exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "ch.ExchangeDeclare")
	}

	q, err := ch.QueueDeclare(
		c.queueName, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return errors.Wrap(err, "ch.QueueDeclare")
	}

	err = ch.QueueBind(
		q.Name,         // queue name
		c.bindingKey,   // routing key
		c.exchangeName, // exchange
		false,
		nil)
	if err != nil {
		return errors.Wrap(err, "ch.QueueBind")
	}

	err = ch.Qos(
		5,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return errors.Wrap(err, "ch.Qos")
	}

	deliveries, err := ch.Consume(
		q.Name,        // queue
		c.consumerTag, // consumer
		false,         // auto ack
		false,         // exclusive
		false,         // no local
		false,         // no wait
		nil,           // args
	)

	if err != nil {
		return errors.Wrap(err, "ch.Consume")
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-deliveries:
				if !ok {
					return
				}
				var integrationEvent event_bus.IntegrationEvent
				err := json.Unmarshal(d.Body, &integrationEvent)
				if err != nil {
					continue
					//return errors.Wrap(err, "ch.Qos")
				} else {
					go handler.Handle(ctx, integrationEvent)
					if err := d.Ack(false); err != nil {

					}
				}
			}
		}
	}()

	// Block until the context is canceled
	<-ctx.Done()
	return nil
}


func handle(deliveries <- chan amqp.Delivery, done chan error) {
	// for d := range deliveries {

	// }
}