package primitive

type ListOption interface {
	Apply(*ListOptions)
}

type ListOptions struct {
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

func (o order) Apply(options *ListOptions) {
	v := o
	options.Order = &v
}

func (o quantity) Apply(options *ListOptions) {
	v := o
	options.Quantity = &v
}
