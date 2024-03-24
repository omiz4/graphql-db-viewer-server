package order

import (
	"context"
	"gqlgen-db-viewer/application/model"
	"time"
)

func (u *OrderUsecase) ListOrders(ctx context.Context, input model.ListRecordsInput) ([]model.Order, error) {
	orders := make([]model.Order, 0)
	orders = append(orders, model.Order{
		Iid:       123,
		Sid:       "hoge",
		UUID:      input.UUID,
		CreatedAt: time.Now(),
	})

	return orders, nil
}
