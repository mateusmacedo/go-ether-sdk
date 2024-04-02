package model

type BaseEntity struct {}

type Entity interface {
	ID() interface{}
}

type entityInstance struct {
	id interface{}
}

type EntityOption func(*entityInstance) error

func NewEntity(opts ...EntityOption) Entity {
	e := &entityInstance{}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *entityInstance) ID() interface{} {
	return e.id
}

func WithID(id interface{}) EntityOption {
	return func(e *entityInstance) error {
		e.id = id
		return nil
	}
}