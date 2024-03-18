package eventbusrabbitmq

import (
	"context"
	event_bus "event-bus/event_bus"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type EventBusJetStream struct {
	js jetstream.JetStream
}

var _ event_bus.EventBus = (*EventBusJetStream)(nil)

func NewEventBusJetStream() (event_bus.EventBus, error) {
	nc, _ := nats.Connect(nats.DefaultURL)

	return &EventBusJetStream{}, nil
}
func (eb *EventBusJetStream) Subcribe(ctx context.Context, handler event_bus.IntegrationEventHandler) error {
	eb.js.PublishMsg(ctx, nats.Msg{
		Subject: "",
		Reply:   "",
		Header:  nil,
		Data:    nil,
		Sub:     nil,
	})
	panic("implement me")
}

func (eb *EventBusJetStream) Unsubscribe(ctx context.Context, handler event_bus.IntegrationEventHandler) error {
	//TODO implement me
	panic("implement me")
}

func (eb *EventBusJetStream) Publish(ctx context.Context, event event_bus.IntegrationEvent) error {
	//TODO implement me
	panic("implement me")
}
