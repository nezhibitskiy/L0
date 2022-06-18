package composite

import (
	"L0/internal"
	"L0/internal/handlers"
)

type OrdersComposite struct {
	h *handlers.Handler
	u internal.Usecase
	r internal.Repository
}

//func NewComposite(pgComposite *connections.PostgresDBComposite) *OrdersComposite {
//	r := repository.NewRepository(pgComposite.DB)
//	u := usecase.NewUsecase(r)
//
//	composite := OrdersComposite{
//		h: nil,
//		u: u,
//		r: r,
//	}
//	return &composite
//}
