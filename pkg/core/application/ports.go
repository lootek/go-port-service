package application

import (
	"context"

	"github.com/lootek/go-port-service/pkg/core/domain"
	"github.com/lootek/go-port-service/pkg/core/services/repository"
)

type Ports struct {
	repo repository.Port
}

func (p *Ports) Run(_ context.Context) {
}

func (p *Ports) Stop() {
}

func NewPorts(r repository.Port) *Ports {
	return &Ports{repo: r}
}

func (p *Ports) List() []domain.Port {
	return p.repo.GetAll()
}

func (p *Ports) Add(port domain.Port) error {
	return p.repo.Insert(port)
}

func (p *Ports) Update(id string, port domain.Port) error {
	return p.Update(id, port)
}
