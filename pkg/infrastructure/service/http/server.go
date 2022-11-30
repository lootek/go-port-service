package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lootek/go-port-service/pkg/core/services/app"
	jsonadapter "github.com/lootek/go-port-service/pkg/infrastructure/adapters/json"
)

type Server struct {
	portService app.PortService
	httpRouter  *gin.Engine
}

func NewServer(srv app.PortService) *Server {
	return &Server{
		portService: srv,
		httpRouter:  gin.New(),
	}
}

func (s *Server) Run(ctx context.Context) {
	s.portService.Run(ctx)

	s.httpRouter.GET("/ports", s.GetAll)
	s.httpRouter.POST("/ports", s.Add)
	s.httpRouter.PUT("/ports/:id", s.Update)
	_ = s.httpRouter.Run("8000") // TODO: Handle this error (and propagate through the Service interface)
}

func (s *Server) Stop() {
	s.portService.Stop()
}

func (s *Server) GetAll(c *gin.Context) {
	ports := s.portService.List()

	result := make([]jsonadapter.Port, len(ports), len(ports))
	for i, p := range ports {
		result[i].FromDomain(p)
	}

	data, err := json.Marshal(result)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if _, err = c.Writer.Write(data); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func (s *Server) Add(c *gin.Context) {
	// TODO
}

func (s *Server) Update(c *gin.Context) {
	// TODO
}
