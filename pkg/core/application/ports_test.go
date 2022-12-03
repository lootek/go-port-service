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

func TestPorts_Upsert(t *testing.T) {
	p := NewPorts(memory.NewStorage())
	p.Run(context.TODO())
	defer p.Stop()

	err := p.Upsert([]domain.Port{
		{
			ID:   "p1_ID",
			Name: "p1_Name",
		},
	})
	require.NoError(t, err)
	require.Len(t, p.List(), 1)
	require.Equal(t, "p1_Name", p.List()[0].Name)

	err = p.Upsert([]domain.Port{
		{
			ID:   "p2_ID",
			Name: "p2_Name",
		},
	})
	require.NoError(t, err)
	require.Len(t, p.List(), 2)
	require.Equal(t, "p1_Name", p.List()[0].Name)
	require.Equal(t, "p2_Name", p.List()[1].Name)

	err = p.Upsert([]domain.Port{
		{
			ID:   "p1_ID",
			Name: "p1_NewName",
		},
	})
	require.NoError(t, err)
	require.Len(t, p.List(), 2)
	require.Equal(t, "p1_NewName", p.List()[0].Name)
	require.Equal(t, "p2_Name", p.List()[1].Name)
}
