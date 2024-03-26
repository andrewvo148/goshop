package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	conn         *amqp.Connection
	exchangeName string
	bindingKey   string
	queueName string
	consumerTag string
}

type Option func(*consumer)

func ExchangeName(exchangeName string) Option {
	return func(c *consumer) {
		c.exchangeName = exchangeName
	}
}

func BindingKey(bindingKey string) Option {
	return func(c *consumer) {
		c.bindingKey = bindingKey
	}
}


func QueueName(queueName string) Option {
	return func(c *consumer) {
		c.queueName = queueName
	}
}

func ConsumerTag(consumerTag string) Option {
	return func (c *consumer)  {
		c.consumerTag = consumerTag
	}
}