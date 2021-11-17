package aws

import (
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/primitive"
)

var _ model.Storage = &storage{}

type storage struct {
}

func NewStorage() model.Storage {
	return storage{}
}

func (s storage) Path() primitive.Path {
	return primitive.Path{}
}

func (s storage) Find() {

}

func (s storage) Remove() {

}
