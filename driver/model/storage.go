package model

import (
	"context"
	"fmt"

	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
	"github.com/t-mutaguchi-10antz/cs-rotate/driver/model/aws"
)

func NewStorage(ctx context.Context, verbose bool, platform string) (model.Storage, error) {
	var storage model.Storage

	p, err := model.Platform(platform)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage struct: %w", err)
	}

	switch p {
	case model.PlatformAWS:
		storage, err = aws.NewStorage(ctx, verbose)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create storage struct: %w", err)
	}

	return storage, nil
}
