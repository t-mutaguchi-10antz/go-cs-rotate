package model

import (
	"context"
)

type Storage interface {
	Rotate(context.Context, ...RotateParam) error
}
