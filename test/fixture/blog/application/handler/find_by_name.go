package handler

import (
	apperr "github.com/mateusmacedo/go-ether-sdk/application/err"
	apphdl "github.com/mateusmacedo/go-ether-sdk/application/handler"
	appmsg "github.com/mateusmacedo/go-ether-sdk/application/message"

	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/application/message"
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

	content := m.Content()
	if len(content) == 0 {
		return nil, apperr.ErrMessageContentEmpty
	}

	if _, ok := m.(*message.FindByNameMessage); !ok {
		return nil, apperr.ErrMessageNotSupported
	}

	if _, err := h.repo.FindByName(string(content)); err != nil {
		return nil, err
	}

	return appmsg.VoidMessage{}, apperr.ErrHandlerNotImplemented
}
