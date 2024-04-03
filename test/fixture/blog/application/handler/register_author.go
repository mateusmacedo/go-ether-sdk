package handler

import (
	"github.com/mateusmacedo/go-ether-sdk/application/err"
	"github.com/mateusmacedo/go-ether-sdk/application/handler"
	"github.com/mateusmacedo/go-ether-sdk/application/message"

	blogmsg "github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/application/message"
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/repository"
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/service"
)

type registerAuthorHandler struct {
	srv  service.RegisterNewAuthorService
	repo repository.AuthorRepository
}

type RegisterAuthorHandlerOption func(*registerAuthorHandler) error

func WithAuthorService(srv service.RegisterNewAuthorService) RegisterAuthorHandlerOption {
	return func(h *registerAuthorHandler) error {
		h.srv = srv
		return nil
	}
}

func WithAuthorRepository(repo repository.AuthorRepository) RegisterAuthorHandlerOption {
	return func(h *registerAuthorHandler) error {
		h.repo = repo
		return nil
	}
}

func NewRegisterAuthorHandler(opts ...RegisterAuthorHandlerOption) handler.Handler {
	h := registerAuthorHandler{}
	for _, opt := range opts {
		opt(&h)
	}
	return &h
}



func (h *registerAuthorHandler) Handle(m message.Message) (message.Message, error){
	if _, ok := m.(*blogmsg.RegisterAuthorMessage); !ok {
		return nil, err.ErrMessageNotSupported
	}
	return nil, err.ErrHandlerNotImplemented
}