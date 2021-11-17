package model

import (
	"fmt"
)

type Object struct {
	URL    url
	Bucket string
	Prefix string
	Key    string
}

func NewObject(url string) (Object, error) {
	o := Object{}

	urlValue, err := NewURL(url)
	if err != nil {
		return o, fmt.Errorf("Failed to create object struct: %w", err)
	}

	o.URL = urlValue

	return o, nil
}
