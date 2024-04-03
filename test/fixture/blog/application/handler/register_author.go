package handler

import (
	"github.com/mateusmacedo/go-ether-sdk/application/err"
	"github.com/mateusmacedo/go-ether-sdk/application/handler"
	"github.com/mateusmacedo/go-ether-sdk/application/message"

	blogmsg "github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/application/message"
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/model"
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/repository"
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/service"
)

type registerAuthorHandler struct {
	srv  service.RegisterNewAuthorService
	repo repository.AuthorRepository
}

type registerAuthorHandlerOption func(*registerAuthorHandler) error

func WithAuthorService(srv service.RegisterNewAuthorService) registerAuthorHandlerOption {
	return func(h *registerAuthorHandler) error {
		h.srv = srv
		return nil
	}
}

func WithAuthorRepository(repo repository.AuthorRepository) registerAuthorHandlerOption {
	return func(h *registerAuthorHandler) error {
		h.repo = repo
		return nil
	}
}

func NewRegisterAuthorHandler(opts ...registerAuthorHandlerOption) handler.Handler {
	h := registerAuthorHandler{}
	for _, opt := range opts {
		opt(&h)
	}
	return &h
}

func (h *registerAuthorHandler) Handle(m message.Message) (message.Message, error) {
	if _, ok := m.(*blogmsg.RegisterAuthorMessage); !ok {
		return nil, err.ErrMessageNotSupported
	}

	if len(m.Content()) == 0{
		return nil, err.ErrMessageContentEmpty
	}

	author := model.NewAuthor(model.WithName(string(m.Content())))

	if err := h.srv.RegisterNewAuthor(author); err != nil {
		return nil, err
	}

	return nil, err.ErrHandlerNotImplemented
}
