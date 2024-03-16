package eventhandling

import (
	"context"
	"basket-api/integration-events/events"
)

type IntegrationEventHandler interface {
	Handle(ctx context.Context, event events.IntegrationEvent) error
}
