package application

import (
	"context"
	"testing"

	"github.com/lootek/go-port-service/pkg/core/domain"
	"github.com/lootek/go-port-service/pkg/infrastructure/repository/memory"
	"github.com/stretchr/testify/require"
)

func TestNewPorts(t *testing.T) {
	p := NewPorts(nil)
	require.NotNil(t, p)

	p = NewPorts(memory.NewStorage())
	require.NotNil(t, p)
}

func TestPorts_Add(t *testing.T) {
	p := NewPorts(memory.NewStorage())
	p.Run(context.TODO())
	defer p.Stop()

	p.Add(domain.Port{
		ID:          "",
		Name:        "",
		City:        "",
		Country:     "",
		Alias:       nil,
		Regions:     nil,
		Coordinates: [2]float64{},
		Province:    "",
		Timezone:    "",
		Unlocs:      nil,
		Code:        "",
	})
}
