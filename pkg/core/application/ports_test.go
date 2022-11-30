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

	err := p.Add(p1)
	require.NoError(t, err)
	require.Len(t, p.List(), 1)
	require.Equal(t, p1, p.List()[0])

	err = p.Add(p2)
	require.NoError(t, err)
	require.Len(t, p.List(), 2)
}

var p1 = domain.Port{
	ID:          "p1_ID",
	Name:        "p1_Name",
	City:        "p1_City",
	Country:     "p1_Country",
	Alias:       []any{"p1_Alias_1", "p1_Alias_1"},
	Regions:     []any{"p1_Region_1", "p1_Region_1"},
	Coordinates: [2]float64{3, 7},
	Province:    "p1_Province",
	Timezone:    "p1_Timezone",
	Unlocs:      []string{"p1_Unlocs_1", "p1_Unlocs_1"},
	Code:        "p1_Code",
}

var p2 = domain.Port{
	ID:          "p2_ID",
	Name:        "p2_Name",
	City:        "p2_City",
	Country:     "p2_Country",
	Alias:       []any{"p2_Alias_1", "p2_Alias_1"},
	Regions:     []any{"p2_Region_1", "p2_Region_1"},
	Coordinates: [2]float64{-2, -5.5},
	Province:    "p2_Province",
	Timezone:    "p2_Timezone",
	Unlocs:      []string{"p2_Unlocs_1", "p2_Unlocs_1"},
	Code:        "p2_Code",
}

func TestPorts_Update(t *testing.T) {
	// TODO
}
