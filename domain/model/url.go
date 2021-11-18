package model

import (
	"fmt"
	net_url "net/url"

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

func (u url) Bucket() (string, error) {
	v, err := net_url.Parse(string(u))
	if err != nil {
		return "", fmt.Errorf("failed to get bucket: %w", err)
	}
	return v.Host, nil
}

func (u url) Prefix() (string, error) {
	v, err := net_url.Parse(string(u))
	if err != nil {
		return "", fmt.Errorf("failed to get bucket: %w", err)
	}
	runes := v.Path
	if string(runes[0]) == "/" {
		runes = runes[1:]
	}
	if string(runes[len(runes)-1]) != "/" {
		runes += "/"
	}
	return runes, nil
}

func (u url) Key() (string, error) {
	return "", nil
}
