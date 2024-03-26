package kafka

import (
	"context"
	event_bus "event-bus/event_bus"
)

type Kafka struct {
}

func (k *Kafka) Subcribe(ctx context.Context, handler event_bus.IntegrationEventHandler) {

}
