package model

import (
	"fmt"

	"github.com/t-mutaguchi-10antz/cs-rotate/validator"
)

type url string

func URL(v string) (url, error) {
	u := url(v)

	if err := validator.CheckValue(&u, "url"); err != nil {
		return u, fmt.Errorf("failed to create URL value: %w", err)
	}

	return u, nil
}

func (u url) Bucket() string {
	return ""
}

func (u url) Prefix() string {
	return ""
}

func (u url) Key() string {
	return ""
}
