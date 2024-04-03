package message

import "github.com/mateusmacedo/go-ether-sdk/application/message"

type FindByNameMessage struct {
	name string
}

func (m FindByNameMessage) Content() message.Content {
	return []byte(m.name)
}

func NewFindByNameMessage(name string) message.Message {
	return &FindByNameMessage{name: name}
}
