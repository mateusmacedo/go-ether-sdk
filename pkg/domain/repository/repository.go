package repository

import (
	"github.com/mateusmacedo/go-ether-sdk/domain/model"
)

type GetByID interface {
	GetByID(id interface{}) (model.Entity, error)
}

type Persist interface {
	Persist(entity model.Entity) error
}
