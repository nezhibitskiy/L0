package repository

import (
	"L0/internal"
	"context"
	"github.com/gomodule/redigo/redis"
	json "github.com/mailru/easyjson"
)

type RedisRepository struct {
	redis *redis.Pool
}

func NewRedisRepository(redis *redis.Pool) internal.RedisRepository {
	return &RedisRepository{redis: redis}
}

func (r *RedisRepository) SaveOrder(ctx context.Context, order *internal.Order) error {
	connRedis := r.redis.Get()
	defer connRedis.Close()

	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}
	_, err = connRedis.Do("SET", order.OrderUid, orderJSON)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) GetOrder(ctx context.Context, uid string) (*internal.Order, error) {
	connRedis := r.redis.Get()
	defer connRedis.Close()

	orderJSON, err := redis.String(connRedis.Do("GET", uid))
	if err != nil {
		return nil, err
	}

	order := internal.Order{}
	err = json.Unmarshal([]byte(orderJSON), &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
