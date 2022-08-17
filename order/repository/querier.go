package repository

import (
	"context"
)

type Querier interface {
	CreateOrder(ctx context.Context, arg CreateOrderParams) error
	GetOrderByID(ctx context.Context, orderUid string) (GetOrderByIDRow, error)
	ListOrders(ctx context.Context, arg ListOrdersParams) ([]ListOrdersRow, error)
}

var _ Querier = (*Queries)(nil)
