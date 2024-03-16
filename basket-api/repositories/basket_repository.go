package repositories

import (
	"basket-api/model"
	"context"
)

type BasketRepository interface {
	GetBasket(ctx context.Context, customerID string) (*model.CustomerBasket, error)
	UpdateBasket(ctx context.Context, basket *model.CustomerBasket) error
	DeleteBasket(ctx context.Context, id string) error
}
