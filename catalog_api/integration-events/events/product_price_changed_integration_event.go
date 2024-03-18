package events

import (
	"event_bus/events"
	"github.com/google/uuid"
	"time"
)

type ProductPriceChangedIntegrationEvent struct {
	events.IntegrationEvent
	ProductID int
	NewPrice  float64
	OldPrice  float64
}

func NewProductPriceChangedIntegrationEvent(productID int, newPrice float64, oldPrice float64) *ProductPriceChangedIntegrationEvent {
	return &ProductPriceChangedIntegrationEvent{
		events.IntegrationEvent{
			ID:           uuid.New(),
			CreationDate: time.Now().UTC(),
		},
		productID,
		newPrice,
		oldPrice,
	}
}
