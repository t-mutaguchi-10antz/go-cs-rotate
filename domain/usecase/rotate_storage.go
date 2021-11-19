package usecase

import (
	"context"
	"fmt"

	"github.com/t-mutaguchi-10antz/go/cs-rotate/domain/model"
)

func (u usecase) RotateStorage(ctx context.Context, url string, quantity uint, order string) error {
	// 外部からの入力を domain primitive な値に変換して安全性を担保する
	urlVal, err := model.WithURL(url)
	if err != nil {
		return fmt.Errorf("failed to rorate storage: %w", err)
	}
	qtyVal, err := model.WithQuantity(quantity)
	if err != nil {
		return fmt.Errorf("failed to rorate storage: %w", err)
	}
	orderVal, err := model.WithOrder(order)
	if err != nil {
		return fmt.Errorf("failed to rorate storage: %w", err)
	}

	// 条件に合わせてクラウドストレージ上のリソースを走査＆削除する
	params := []model.RotateParam{urlVal, qtyVal, orderVal}
	err = u.Model().Storage().Rotate(ctx, params...)
	if err != nil {
		return fmt.Errorf("failed to rorate storage: %w", err)
	}

	return nil
}
