package events

import (
	"time"

	"github.com/google/uuid"
)

type OrderStartedIntegrationEvent struct {
	IntegrationEvent
	UserID string
}

func NewOrderStartedIntegrationEvent(userID string) *OrderStartedIntegrationEvent {
	return &OrderStartedIntegrationEvent{
		IntegrationEvent: IntegrationEvent{
			ID:           uuid.New(),
			CreationDate: time.Now().UTC(),
		},
		UserID: userID,
	}

}
