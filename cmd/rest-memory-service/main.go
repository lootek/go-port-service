package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/lootek/go-port-service/pkg/core/application"
	"github.com/lootek/go-port-service/pkg/infrastructure/repository/memory"
	"github.com/lootek/go-port-service/pkg/infrastructure/service/http"
)

func main() {
	ctx, cancelFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancelFn()

	portsRepo := memory.NewStorage(ctx)
	portsApp := application.NewPorts(portsRepo)
	httpSrv := http.NewServer(portsApp)

	httpSrv.Run(ctx)

	select {
	case <-ctx.Done():
		httpSrv.Stop()
	}
}
