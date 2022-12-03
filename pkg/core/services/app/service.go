package app

import (
	"context"

	"github.com/lootek/go-port-service/pkg/core/domain"
)

type PortService interface {
	Service
	PortRepository
}

type Service interface {
	Run(ctx context.Context)
	Stop()
}

type PortRepository interface {
	List() ([]domain.Port, error)
	Add(port domain.Port) (string, error)
	Update(id string, port domain.Port) error
}
