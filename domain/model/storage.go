package model

import (
	"context"
)

type Storage interface {
	List(context.Context, ...ListOption) (objects []Object, nextOptions []ListOption, err error)
	Delete(context.Context, []Object) error
}
