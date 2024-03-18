package eventbusrabbitmq

import (
	"context"
	events "event-bus/event_bus"

	"github.com/nats-io/nats.go/jetstream"
)

type EventBusJetStream struct {
	js jetstream.JetStream
}


var _ EventBus = (*EventBusJetStream)(nil)



// func NewEventBusJetStream(js jetstream.JetStream) *EventBusJetStream {
// 	return &EventBusJetStream{}
// }

// func (eb *EventBusJetStream) Publish(ctx context.Context, subject, evt events.IntegrationEvent) error {
// 	eb.js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{

// 	})

	
// 	eb.js.PublishAsync(subject)
// }
