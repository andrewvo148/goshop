package repositories

import (
	"basket-api/model"
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisBasketRepository struct {
	logger   *log.Logger
	database *redis.Client
}

// NewRedisBasketRepository creates a new instance of RedisBasketRepository.
func NewRedisBasketRepository(logger *log.Logger, redis *redis.Client) *RedisBasketRepository {
	return &RedisBasketRepository{
		logger:   logger,
		database: redis,
	}
}

func (r *RedisBasketRepository) GetBasket(ctx context.Context, customerID string) (*model.CustomerBasket, error) {
	data, err := r.database.Get(ctx, GetBasketKey(customerID)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var basket model.CustomerBasket
	err = json.Unmarshal(data, &basket)
	if err != nil {
		return nil, err
	}

	return &basket, nil
}

func (r *RedisBasketRepository) UpdateBasket(ctx context.Context, basket *model.CustomerBasket) error {
	jsonBytes, err := json.Marshal(basket)
	if err != nil {
		return err
	}

	err = r.database.Set(ctx, GetBasketKey(basket.BuyerID), jsonBytes, 0).Err()
	if err != nil {
		r.logger.Printf("Problem occurred persisting the item: %v", err)
		return err
	}

	r.logger.Println("Basket item persisted successfully.")
	return nil
}

func (r *RedisBasketRepository) DeleteBasket(ctx context.Context, id string) error {
	return r.database.Del(ctx, GetBasketKey(id)).Err()
}

func GetBasketKey(userID string) string {
	return "/basket/" + userID
}
