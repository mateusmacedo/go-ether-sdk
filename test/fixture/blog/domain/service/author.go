package service

import (
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/model"
)

type RegisterNewAuthorService interface {
	RegisterNewAuthor(author *model.Author) error
}
