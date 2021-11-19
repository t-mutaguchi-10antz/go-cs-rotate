package model

type RotateParam interface {
	Apply(*RotateParams)
}

type RotateParams struct {
	URL      url
	Order    order
	Quantity quantity
}

func NewRotateParams(params ...RotateParam) RotateParams {
	opts := RotateParams{}
	for _, param := range params {
		if param == nil {
			continue
		}
		param.Apply(&opts)
	}
	return opts
}

func (o url) Apply(params *RotateParams) {
	params.URL = o
}

func (o order) Apply(params *RotateParams) {
	params.Order = o
}

func (o quantity) Apply(params *RotateParams) {
	params.Quantity = o
}
