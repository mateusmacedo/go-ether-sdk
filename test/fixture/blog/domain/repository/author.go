package repository

import (
	apprepo "github.com/mateusmacedo/go-ether-sdk/domain/repository"

	"github.com/mateusmacedo/go-ether-sdk/test/fixture/blog/domain/model"
)

type FindByNameResult []*model.Author

type FindByName interface {
	FindByName(name string) (FindByNameResult, error)
}

type AuthorRepository interface {
	apprepo.GetByID
	apprepo.Persist
	FindByName
}