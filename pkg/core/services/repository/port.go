package repository

import (
	"github.com/lootek/go-port-service/pkg/core/domain"
)

type Port interface {
	GetAll() []domain.Port
	Insert(port domain.Port) error
	Update(id string, port domain.Port) error
}
