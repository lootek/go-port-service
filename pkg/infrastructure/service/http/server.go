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
	"github.com/lootek/go-port-service/pkg/core/services/app"
	jsonadapter "github.com/lootek/go-port-service/pkg/infrastructure/adapters/json"
)

type Server struct {
	portService app.PortService
	httpRouter  *gin.Engine
	srvShutdown func()
}

func NewServer(srv app.PortService) *Server {
	return &Server{
		portService: srv,
		httpRouter:  gin.New(),
	}
}

func (s *Server) Run(ctx context.Context) {
	s.portService.Run(ctx)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: s.httpRouter,
	}
	s.srvShutdown = func() {
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf("HTTP server shutdown failed: %s\n", err)
		}
	}

	s.httpRouter.MaxMultipartMemory = 16 << 20 // 16 MiB
	s.httpRouter.GET("/ports", s.GetAll)
	s.httpRouter.POST("/ports", s.Add)
	s.httpRouter.PUT("/ports/:id", s.Update)

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
	ports, _ := s.portService.List()

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
	// file, err := c.FormFile("ports")
	// if err != nil {
	// 	_ = c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 32<<20+1024)
	reader, err := c.Request.MultipartReader()
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	p, _ := reader.NextPart()
	for {
		if p.FormName() == "file_field" {
			break
		}
		p, _ = reader.NextPart()
	}

	buf := bufio.NewReader(p)
	sniff, _ := buf.Peek(512)
	contentType := http.DetectContentType(sniff)
	if contentType != "application/json" {
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("file type not allowed"))
			return
		}

		var maxSize int64 = 32 << 20
		mr := io.MultiReader(buf, io.LimitReader(p, maxSize-511))
		dec := json.NewDecoder(mr)

		var ports []jsonadapter.Port
		err := dec.Decode(&ports)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func (s *Server) Update(c *gin.Context) {
	// TODO
}
