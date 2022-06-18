package internal

import "context"

type Repository interface {
	GetOrders(ctx context.Context) ([]Order, error)
	GetOrder(ctx context.Context, uid string) (*Order, error)
	SaveOrder(ctx context.Context, data *Order) error
}
type RedisRepository interface {
	GetOrder(ctx context.Context, uid string) (*Order, error)
	SaveOrder(ctx context.Context, data *Order) error
}
