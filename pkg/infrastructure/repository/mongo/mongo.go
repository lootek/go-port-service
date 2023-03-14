package mongo

import (
	"context"

	"github.com/lootek/go-port-service/pkg/core/domain"
)

type DB struct {
}

func NewDB(_ context.Context) *DB {
	return &DB{}
}

func (D *DB) GetAll() ([]domain.Port, error) {
	// TODO implement me
	// Remember to create the new ctx for a db call request deriving from the main ctx
	panic("implement me")
}

func (D *DB) Insert(p domain.Port) (string, error) {
	// TODO implement me
	// Remember to create the new ctx for a db call request deriving from the main ctx
	panic("implement me")
}

func (D *DB) Update(id string, p domain.Port) error {
	// TODO implement me
	// Remember to create the new ctx for a db call request deriving from the main ctx
	panic("implement me")
}
