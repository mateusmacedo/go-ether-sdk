package helper

import "github.com/mateusmacedo/go-ether-sdk/application/message"

type InvalidMessage struct{
}

func (m *InvalidMessage) Content() message.Content {
	return []byte("invalid")
}