package repository

import (
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/model"
)

type FindByNameResult []*model.Author

type FindByName interface {
	FindByName(name string) (FindByNameResult, error)
}

type AuthorRepository interface {
	FindByName
	Persist(author *model.Author) error
}