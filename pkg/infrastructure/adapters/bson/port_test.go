package bson

import (
	"reflect"
	"testing"

	"github.com/lootek/go-port-service/pkg/core/domain"
)

func TestPort_FromDomain(t *testing.T) {
	type fields struct {
		ID          string
		Name        string
		City        string
		Country     string
		Alias       []any
		Regions     []any
		Coordinates [2]float64
		Province    string
		Timezone    string
		Unlocs      []string
		Code        string
	}
	type args struct {
		dp domain.Port
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Port{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				City:        tt.fields.City,
				Country:     tt.fields.Country,
				Alias:       tt.fields.Alias,
				Regions:     tt.fields.Regions,
				Coordinates: tt.fields.Coordinates,
				Province:    tt.fields.Province,
				Timezone:    tt.fields.Timezone,
				Unlocs:      tt.fields.Unlocs,
				Code:        tt.fields.Code,
			}
			p.FromDomain(tt.args.dp)
		})
	}
}

func TestPort_ToDomain(t *testing.T) {
	type fields struct {
		ID          string
		Name        string
		City        string
		Country     string
		Alias       []any
		Regions     []any
		Coordinates [2]float64
		Province    string
		Timezone    string
		Unlocs      []string
		Code        string
	}
	tests := []struct {
		name   string
		fields fields
		want   domain.Port
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Port{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				City:        tt.fields.City,
				Country:     tt.fields.Country,
				Alias:       tt.fields.Alias,
				Regions:     tt.fields.Regions,
				Coordinates: tt.fields.Coordinates,
				Province:    tt.fields.Province,
				Timezone:    tt.fields.Timezone,
				Unlocs:      tt.fields.Unlocs,
				Code:        tt.fields.Code,
			}
			if got := p.ToDomain(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
