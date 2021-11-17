package model

import (
	"fmt"

	"github.com/t-mutaguchi-10antz/cs-rotate/validator"
)

type URL struct {
	Value    string `validate:"url"`
	Protocol string
	Bucket   string
	Key      string
}

func NewURL(v string) (URL, error) {
	u := URL{}

	if err := validator.CheckStruct(&u); err != nil {
		return u, fmt.Errorf("Failed to create URL struct: %w", err)
	}

	u.Value = v

	return u, nil
}
