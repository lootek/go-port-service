package mongo

import (
	"github.com/lootek/go-port-service/pkg/core/domain"
)

type DB struct {
}

func NewDB() *DB {
	return &DB{}
}

func (D DB) GetAll() []domain.Port {
	// TODO implement me
	panic("implement me")
}

func (D DB) Insert(port domain.Port) error {
	// TODO implement me
	panic("implement me")
}

func (D DB) Update(id string, port domain.Port) error {
	// TODO implement me
	panic("implement me")
}
