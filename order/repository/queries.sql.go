package repository

import (
	"context"
	"encoding/json"
)

const createOrder = `-- name: CreateOrder :exec
INSERT INTO orders(
    order_uid, data
) VALUES (
    $1, $2
)
`

type CreateOrderParams struct {
	OrderUid string          `json:"order_uid"`
	Data     json.RawMessage `json:"data"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) error {
	_, err := q.db.ExecContext(ctx, createOrder, arg.OrderUid, arg.Data)
	return err
}

const getOrderByID = `-- name: GetOrderByID :one
SELECT order_uid, data
FROM orders
WHERE order_uid = $1
`

type GetOrderByIDRow struct {
	OrderUid string          `json:"order_uid"`
	Data     json.RawMessage `json:"data"`
}

func (q *Queries) GetOrderByID(ctx context.Context, orderUid string) (GetOrderByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getOrderByID, orderUid)
	var i GetOrderByIDRow
	err := row.Scan(&i.OrderUid, &i.Data)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT order_uid, data
FROM orders
WHERE id > $1
LIMIT $2
`

type ListOrdersParams struct {
	ID    int64 `json:"id"`
	Limit int32 `json:"limit"`
}

type ListOrdersRow struct {
	OrderUid string          `json:"order_uid"`
	Data     json.RawMessage `json:"data"`
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]ListOrdersRow, error) {
	rows, err := q.db.QueryContext(ctx, listOrders, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListOrdersRow
	for rows.Next() {
		var i ListOrdersRow
		if err := rows.Scan(&i.OrderUid, &i.Data); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
