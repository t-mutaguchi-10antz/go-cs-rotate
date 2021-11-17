package model

type Model interface {
	Storage() Storage
}

var _ Model = &model{}

type model struct {
	storage Storage
}

func NewModel(storage Storage) Model {
	return model{
		storage: storage,
	}
}

func (m model) Storage() Storage {
	return m.storage
}
