package usecase

import (
	"context"
	"fmt"

	"github.com/t-mutaguchi-10antz/cs-rotate/domain/primitive"
)

func (u usecase) RotateStorage(ctx context.Context, order string, quantity uint) error {
	o, err := primitive.NewOrder(order)
	if err != nil {
		return fmt.Errorf("Failed to rorate storage: %w", err)
	}

	q, err := primitive.NewQuantity(quantity)
	if err != nil {
		return fmt.Errorf("Failed to rorate storage: %w", err)
	}

	resources, err := u.Model().Storage().List(ctx, o, q)
	if err != nil {
		return fmt.Errorf("Failed to rorate storage: %w", err)
	}

	err = u.Model().Storage().Delete(ctx, resources)
	if err != nil {
		return fmt.Errorf("Failed to rorate storage: %w", err)
	}

	return nil
}
