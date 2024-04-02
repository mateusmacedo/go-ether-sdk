package model

type Author struct {
	id      int
	version int
	name    string
}

type AuthorOption func(*Author) error

func NewAuthor(opts ...AuthorOption) *Author {
	a := &Author{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func (a *Author) ID() int {
	return a.id
}

func (a *Author) Version() int {
	return a.version
}

func (a *Author) Name() string {
	return a.name
}

func WithID(id int) AuthorOption {
	return func(a *Author) error {
		a.id = id
		return nil
	}
}

func WithVersion(version int) AuthorOption {
	return func(a *Author) error {
		a.version = version
		return nil
	}
}

func WithName(name string) AuthorOption {
	return func(a *Author) error {
		a.name = name
		return nil
	}
}
