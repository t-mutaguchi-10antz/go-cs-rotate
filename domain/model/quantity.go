package model

import (
	"fmt"

	"github.com/t-mutaguchi-10antz/go/cs-rotate/validator"
)

type quantity uint

func WithQuantity(v uint) (quantity, error) {
	if err := validator.CheckValue(v, "min=1"); err != nil {
		return quantity(0), fmt.Errorf("failed to create quantity value: %w", err)
	}
	return quantity(v), nil
}
