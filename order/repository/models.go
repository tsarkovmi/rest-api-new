package repository

import (
	"encoding/json"
)

type Order struct {
	ID       int64           `json:"id"`
	OrderUid string          `json:"order_uid"`
	Data     json.RawMessage `json:"data"`
}
