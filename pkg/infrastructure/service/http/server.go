package http

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lootek/go-port-service/pkg/core/domain"
	"github.com/lootek/go-port-service/pkg/core/services/app"
	jsonadapter "github.com/lootek/go-port-service/pkg/infrastructure/adapters/json"
)

type Server struct {
	portService app.PortService
	httpRouter  *gin.Engine
	srvShutdown func()
}

func NewServer(portsApp app.PortService) *Server {
	s := &Server{
		portService: portsApp,
		httpRouter:  gin.New(),
	}

	s.httpRouter.MaxMultipartMemory = 16 << 20 // 16 MiB
	s.httpRouter.GET("/ports", s.GetAll)
	s.httpRouter.POST("/ports", s.Upsert)

	return s
}

func (s *Server) Run(ctx context.Context) {
	s.portService.Run(ctx)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: s.httpRouter,

		// TODO: Add TLS support
		TLSConfig: nil,
	}
	s.srvShutdown = func() {
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf("HTTP server shutdown failed: %s\n", err)
		}
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("HTTP server error: %s\n", err)
		}
	}()
}

func (s *Server) Stop() {
	s.srvShutdown()
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

func (s *Server) Upsert(c *gin.Context) {
	fh, err := c.FormFile("ports")
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	file, err := fh.Open()
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	buf := bufio.NewReader(file)
	sniff, err := buf.Peek(512)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	contentType := http.DetectContentType(sniff)
	if contentType != "application/json" {
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("file type not allowed"))
			return
		}

		var maxSize int64 = 32 << 20
		mr := io.MultiReader(buf, io.LimitReader(file, maxSize-511))
		dec := json.NewDecoder(mr)

		var portsJSON []jsonadapter.Port
		err := dec.Decode(&portsJSON)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ports := make([]domain.Port, 0, len(portsJSON))
		for _, pj := range portsJSON {
			ports = append(ports, pj.ToDomain())
		}

		err = s.portService.Upsert(ports)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	}
}
