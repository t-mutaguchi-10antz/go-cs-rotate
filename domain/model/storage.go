package model

import (
	"context"
)

type Storage interface {
	List(context.Context) ([]Resource, error)
	Delete(context.Context, []Resource) error
}
