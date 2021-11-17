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
	urlValue, err := URL(url)
	if err != nil {
		return Object{}, fmt.Errorf("Failed to create object struct: %w", err)
	}

	return Object{
		URL:    urlValue,
		Bucket: urlValue.Bucket(),
		Prefix: urlValue.Prefix(),
		Key:    urlValue.Key(),
	}, nil
}
