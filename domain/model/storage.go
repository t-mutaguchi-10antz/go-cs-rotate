package model

import (
	"context"
)

type Storage interface {
	List(context.Context, ...ListOption) ([]Object, error)
	Delete(context.Context, []Object) error
}
