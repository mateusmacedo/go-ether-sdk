package handler

import (
	apperr "github.com/mateusmacedo/go-ether-sdk/application/err"
	apphdl "github.com/mateusmacedo/go-ether-sdk/application/handler"
	appmsg "github.com/mateusmacedo/go-ether-sdk/application/message"

	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/repository"
)

type findByNameHandler struct {
	repo repository.FindByName
}

type findByNameHandlerOption func(*findByNameHandler) error

func WithFindByNameRepository(repo repository.FindByName) findByNameHandlerOption {
	return func(h *findByNameHandler) error {
		h.repo = repo
		return nil
	}
}

func NewFindByNameHandler(opts ...findByNameHandlerOption) apphdl.Handler {
	h := findByNameHandler{}
	for _, opt := range opts {
		opt(&h)
	}
	return &h
}

func (h *findByNameHandler) Handle(m appmsg.Message) (appmsg.Message, error) {
	if m == nil {
		return nil, apperr.ErrMessageContentEmpty
	}

	return appmsg.VoidMessage{}, apperr.ErrHandlerNotImplemented
}