package event_bus

import "context"

type EventBus interface {
	Subcribe(ctx context.Context, handler IntegrationEventHandler) error
	Unsubscribe(ctx context.Context, handler IntegrationEventHandler) error
	Publish(ctx context.Context, event IntegrationEvent) error
}
