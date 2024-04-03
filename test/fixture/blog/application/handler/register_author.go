package handler

import (
	apperr "github.com/mateusmacedo/go-ether-sdk/application/err"
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
	content := m.Content()
	if len(content) == 0{
		return nil, apperr.ErrMessageContentEmpty
	}

	if _, ok := m.(*blogmsg.RegisterAuthorMessage); !ok {
		return nil, apperr.ErrMessageNotSupported
	}

	name := string(content)

	if got, err := h.repo.FindByName(name); got != nil {
		return nil, apperr.ErrPersistenceConflict
	} else if err != nil {
		return nil, err
	}

	author := model.NewAuthor(model.WithName(name))

	if err := h.srv.RegisterNewAuthor(author); err != nil {
		return nil, err
	}

	if err := h.repo.Persist(author); err != nil {
		return nil, err
	}

	return message.VoidMessage{}, nil
}
