package model

import (
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/primitive"
)

type Storage interface {
	Path() primitive.Path
	Find()
	Remove()
}
