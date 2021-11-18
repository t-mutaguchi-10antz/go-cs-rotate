package model

import (
	"fmt"

	"github.com/t-mutaguchi-10antz/cs-rotate/validator"
)

type url string

func WithURL(v string) (url, error) {
	u := url(v)

	if err := validator.CheckValue(&u, "url"); err != nil {
		return u, fmt.Errorf("failed to create URL value: %w", err)
	}

	return u, nil
}

func (u url) Bucket() string {
	return "resource-dev1.game.prince-royale.jp"
}

func (u url) Prefix() string {
	return "AssetBundle/"
}

func (u url) Key() string {
	return ""
}
