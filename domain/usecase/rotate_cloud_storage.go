package usecase

import (
	"context"
	"fmt"

	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
)

func (u usecase) RotateCloudStorage(ctx context.Context, url string, quantity uint, order string) error {
	// 外部からの入力を domain primitive な値に変換して安全性を担保する
	urlVal, err := model.URL(url)
	if err != nil {
		return fmt.Errorf("failed to rorate storage: %w", err)
	}
	qtyVal, err := model.Quantity(quantity)
	if err != nil {
		return fmt.Errorf("failed to rorate storage: %w", err)
	}
	orderVal, err := model.Order(order)
	if err != nil {
		return fmt.Errorf("failed to rorate storage: %w", err)
	}

	// 条件に合致するクラウドストレージ上のリソースを走査・削除する
	var objects []model.Object
	options := []model.ListOption{urlVal, qtyVal, orderVal}
	for options != nil {
		objects, options, err = u.Model().Storage().List(ctx, options...)
		if err != nil {
			return fmt.Errorf("failed to rorate storage: %w", err)
		}
		err = u.Model().Storage().Delete(ctx, objects)
		if err != nil {
			return fmt.Errorf("failed to rorate storage: %w", err)
		}
	}

	return nil
}
