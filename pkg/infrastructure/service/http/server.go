package http

import (
	"context"

	"github.com/lootek/go-port-service/pkg/core/services/app"
)

type Server struct {
	portService app.PortService
}

func NewServer(srv app.PortService) *Server {
	return &Server{
		portService: srv,
	}
}

func (s Server) Run(ctx context.Context) {
	s.portService.Run(ctx)
}

func (s Server) Stop() {
	s.portService.Stop()
}
