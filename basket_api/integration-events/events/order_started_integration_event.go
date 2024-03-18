package events

import (
	"event_bus/events"
	"time"

	"github.com/google/uuid"
)

type OrderStartedIntegrationEvent struct {
	events.IntegrationEvent
	UserID string
}

func NewOrderStartedIntegrationEvent(userID string) *OrderStartedIntegrationEvent {
	return &OrderStartedIntegrationEvent{
		IntegrationEvent: events.IntegrationEvent{
			ID:           uuid.New(),
			CreationDate: time.Now().UTC(),
		},
		UserID: userID,
	}

}
