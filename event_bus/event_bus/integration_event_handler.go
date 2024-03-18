package event_bus

import (
	"context"
)

type IntegrationEventHandler interface {
	Handle(ctx context.Context, event IntegrationEvent) error
}
