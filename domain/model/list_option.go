package model

type ListOption interface {
	Apply(*ListOptions)
}

type ListOptions struct {
	URL      *url
	Order    *order
	Quantity *quantity
}

func NewListOptions(options ...ListOption) ListOptions {
	opts := ListOptions{}
	for _, option := range options {
		if option == nil {
			continue
		}
		option.Apply(&opts)
	}
	return opts
}

func (o url) Apply(options *ListOptions) {
	v := o
	options.URL = &v
}

func (o order) Apply(options *ListOptions) {
	v := o
	options.Order = &v
}

func (o quantity) Apply(options *ListOptions) {
	v := o
	options.Quantity = &v
}
