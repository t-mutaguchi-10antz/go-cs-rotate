package model

import (
	"fmt"

	"github.com/t-mutaguchi-10antz/cs-rotate/validator"
)

type url string

func NewURL(v string) (url, error) {
	u := url(v)

	if err := validator.CheckValue(&u, "url"); err != nil {
		return u, fmt.Errorf("Failed to create URL value: %w", err)
	}

	return u, nil
}
