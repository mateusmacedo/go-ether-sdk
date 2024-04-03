package repository

import (
	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/model"
)

type FindByName interface {
	FindByName(name string) (*model.Author, error)
}

type AuthorRepository interface {
	FindByName
	Persist(author *model.Author) error
}