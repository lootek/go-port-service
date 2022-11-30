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

func (p *Port) FromDomain(dp domain.Port) {
	p = &Port{
		ID:          dp.ID,
		Name:        dp.Name,
		City:        dp.City,
		Country:     dp.Country,
		Alias:       dp.Alias,
		Regions:     dp.Regions,
		Coordinates: dp.Coordinates,
		Province:    dp.Province,
		Timezone:    dp.Timezone,
		Unlocs:      dp.Unlocs,
		Code:        dp.Code,
	}
}

func (p *Port) ToDomain() domain.Port {
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
