package model

import (
	"fmt"
)

type order string

const (
	OrderAsc  = order("asc")
	OrderDesc = order("desc")
)

func WithOrder(v string) (order, error) {
	o := order(v)
	switch o {
	case OrderAsc, OrderDesc:
		return o, nil
	default:
		return order(""), fmt.Errorf("failed to create order value: invalid value")
	}
}
