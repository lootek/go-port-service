package json

import (
	"github.com/lootek/go-port-service/pkg/core/domain"
)

type Port struct {
	ID          string     `json:"ID,omitempty"`
	Name        string     `json:"name,omitempty"`
	City        string     `json:"city,omitempty"`
	Country     string     `json:"country,omitempty"`
	Alias       []any      `json:"alias,omitempty"`
	Regions     []any      `json:"regions,omitempty"`
	Coordinates [2]float64 `json:"coordinates,omitempty"`
	Province    string     `json:"province,omitempty"`
	Timezone    string     `json:"timezone,omitempty"`
	Unlocs      []string   `json:"unlocs,omitempty"`
	Code        string     `json:"code,omitempty"`
}

func ToJSON(p domain.Port) Port {
	return Port{
		ID:          p.ID,
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}
}

func FromJSON(p Port) domain.Port {
	return domain.Port{
		ID:          p.ID,
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}
}
