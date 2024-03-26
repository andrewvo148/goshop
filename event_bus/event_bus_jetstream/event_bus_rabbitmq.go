package eventbusrabbitmq

import (
	"context"
	event_bus "event-bus/event_bus"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog"
	"sync"
)

type EventBusJetStream struct {
	js     jetstream.JetStream
	logger zerolog.Logger
	mut    sync.Mutex
}

var _ event_bus.EventBus = (*EventBusJetStream)(nil)

func NewEventBusJetStream(js jetstream.JetStream, logger zerolog.Logger) event_bus.EventBus {
	return &EventBusJetStream{
		js:     js,
		logger: logger,
	}
}
func (eb *EventBusJetStream) Subcribe(ctx context.Context, handler event_bus.IntegrationEventHandler) error {
	cfg := nats.ConsumerConfig{
		Durable:            "",
		Name:               "",
		Description:        "",
		DeliverPolicy:      0,
		OptStartSeq:        0,
		OptStartTime:       nil,
		AckPolicy:          0,
		AckWait:            0,
		MaxDeliver:         0,
		BackOff:            nil,
		FilterSubject:      "",
		FilterSubjects:     nil,
		ReplayPolicy:       0,
		RateLimit:          0,
		SampleFrequency:    "",
		MaxWaiting:         0,
		MaxAckPending:      0,
		FlowControl:        false,
		Heartbeat:          0,
		HeadersOnly:        false,
		MaxRequestBatch:    0,
		MaxRequestExpires:  0,
		MaxRequestMaxBytes: 0,
		DeliverSubject:     "",
		DeliverGroup:       "",
		InactiveThreshold:  0,
		Replicas:           0,
		MemoryStorage:      false,
		Metadata:           nil,
	}
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
