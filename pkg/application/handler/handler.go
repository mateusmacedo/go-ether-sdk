package handler

import "github.com/mateusmacedo/go-ether-sdk/application/message"


type Handler interface {
	Handle(m message.Message) (message.Message, error)
}