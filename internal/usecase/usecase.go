package usecase

import (
	"L0/internal"
	"context"
	"log"
)

type Usecase struct {
	r     internal.Repository
	redis internal.RedisRepository
}

func NewUsecase(r internal.Repository, redis internal.RedisRepository) internal.Usecase {
	return &Usecase{r: r, redis: redis}
}

func (u *Usecase) SaveData(ctx context.Context, data *internal.Order) error {
	// TODO: Implement data saving in postgres and redis
	// Handle the message
	log.Printf("Subscribed message for Order: %+v\n", data.OrderUid)
	err := u.r.SaveOrder(ctx, data)
	if err != nil {
		return err
	}
	err = u.redis.SaveOrder(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) InitCache(ctx context.Context) error {
	// TODO: Implement cache filling from postgres
	go func() {
		orders, err := u.r.GetOrders(ctx)
		if err != nil {
			// TODO: Add logger
			return
		}
		for _, order := range orders {
			err = u.redis.SaveOrder(ctx, &order)
			if err != nil {
				// TODO: Add logger
				return
			}
		}
	}()
	return nil
}
func (u *Usecase) GetOrder(ctx context.Context, uid string) (*internal.Order, error) {
	order, err := u.redis.GetOrder(ctx, uid)
	if err != nil {
		return nil, err
	}
	if order != nil {
		return order, nil
	}
	order, err = u.r.GetOrder(ctx, uid)
	if err != nil {
		return nil, err
	}
	return order, nil
}
