package message

import "github.com/mateusmacedo/go-ether-sdk/application/message"

type RegisterAuthorMessage struct {
	name string
}

type RegisterAuthorOption func(*RegisterAuthorMessage) error

func WithName(name string) RegisterAuthorOption {
	return func(m *RegisterAuthorMessage) error {
		m.name = name
		return nil
	}
}

func NewRegisterAuthor(opts ...RegisterAuthorOption) message.Message {
	m := &RegisterAuthorMessage{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m RegisterAuthorMessage) Content() message.Content {
	return []byte(m.name)
}