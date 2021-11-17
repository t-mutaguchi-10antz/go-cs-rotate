package model

import (
	"context"

	"github.com/t-mutaguchi-10antz/cs-rotate/domain/primitive"
)

type Storage interface {
	List(context.Context, ...primitive.ListOption) ([]Resource, error)
	Delete(context.Context, []Resource) error
}
