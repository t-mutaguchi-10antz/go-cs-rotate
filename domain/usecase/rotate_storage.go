package usecase

import (
	"context"
	"fmt"

	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
)

func (u usecase) RotateStorage(ctx context.Context, url string, quantity uint, order string) error {
	urlValue, err := model.URL(url)
	if err != nil {
		return fmt.Errorf("Failed to rorate storage: %w", err)
	}

	qtyValue, err := model.Quantity(quantity)
	if err != nil {
		return fmt.Errorf("Failed to rorate storage: %w", err)
	}

	orderValue, err := model.Order(order)
	if err != nil {
		return fmt.Errorf("Failed to rorate storage: %w", err)
	}

	resources, err := u.Model().Storage().List(ctx, urlValue, qtyValue, orderValue)
	if err != nil {
		return fmt.Errorf("Failed to rorate storage: %w", err)
	}

	err = u.Model().Storage().Delete(ctx, resources)
	if err != nil {
		return fmt.Errorf("Failed to rorate storage: %w", err)
	}

	return nil
}
