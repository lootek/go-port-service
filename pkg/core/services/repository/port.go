package repository

import (
	"github.com/lootek/go-port-service/pkg/core/domain"
)

type Port interface {
	GetAll() ([]domain.Port, error)
	Insert(p domain.Port) (string, error)
	Update(id string, p domain.Port) error
}
