package model

import (
	"fmt"
)

type Resource struct {
	URL URL
}

func NewResource(url string) (Resource, error) {
	r := Resource{}

	u, err := NewURL(url)
	if err != nil {
		return r, fmt.Errorf("Failed to create resource struct: %w", err)
	}

	r.URL = u

	return r, nil
}
