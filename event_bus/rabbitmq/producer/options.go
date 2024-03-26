package producer

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type producer struct {
	conn         *amqp.Connection
	exchangeName string
	bindingKey   string
}

type Option func(*producer)

func ExchangeName(exchangeName string) Option {
	return func(p *producer) {
		p.exchangeName = exchangeName
	}
}

func BindingKey(bindingKey string) Option {
	return func(p *producer) {
		p.bindingKey = bindingKey
	}
}
