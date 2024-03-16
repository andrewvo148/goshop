package events

import (
	"time"

	"github.com/google/uuid"
)

type IntegrationEvent struct {
	ID           uuid.UUID `json:"id"`
	CreationDate time.Time `json:"creation_date"`
}

func NewIntegrationEvent() *IntegrationEvent {
	return &IntegrationEvent{
		ID:           uuid.New(),
		CreationDate: time.Now().UTC(),
	}
}
