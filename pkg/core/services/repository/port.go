package repository

import (
	"github.com/lootek/go-port-service/pkg/core/domain"
)

type Port interface {
	GetAll() []domain.Port
	Upsert(port []domain.Port) error
}
