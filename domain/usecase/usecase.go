package usecase

import (
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
)

type Usecase interface {
	Model() model.Model
}

var _ Usecase = &usecase{}

type usecase struct {
	model model.Model
}

func (u usecase) Model() model.Model {
	return u.model
}

func NewUsecase(model model.Model) usecase {
	return usecase{
		model: model,
	}
}
