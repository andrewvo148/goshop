package eventhandling

import (
	"basket-api/integration-events/events"
	"basket-api/repositories"
	"context"

	"github.com/rs/zerolog/log"
)

type OrderStartedIntegrationEventHandler struct {
	repository repositories.BasketRepository
}

func NewOrderStartedIntegrationEventHandler(repository repositories.BasketRepository) *OrderStartedIntegrationEventHandler {
	return &OrderStartedIntegrationEventHandler{repository: repository}
}

func (h *OrderStartedIntegrationEventHandler) Handle(ctx context.Context, event events.OrderStartedIntegrationEvent) error {
	log.Info().Msgf("Handling integration event: %s - (%+v)", event.ID, event)

	return h.repository.DeleteBasket(ctx, event.UserID)
}
