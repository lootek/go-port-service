package bson

import (
	"github.com/lootek/go-port-service/pkg/core/domain"
)

type Port struct {
	ID          string     `bson:"ID,omitempty"`
	Name        string     `bson:"name,omitempty"`
	City        string     `bson:"city,omitempty"`
	Country     string     `bson:"country,omitempty"`
	Alias       []any      `bson:"alias,omitempty"`
	Regions     []any      `bson:"regions,omitempty"`
	Coordinates [2]float64 `bson:"coordinates,omitempty"`
	Province    string     `bson:"province,omitempty"`
	Timezone    string     `bson:"timezone,omitempty"`
	Unlocs      []string   `bson:"unlocs,omitempty"`
	Code        string     `bson:"code,omitempty"`
}

func ToBSON(p domain.Port) Port {
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

func FromBSON(p Port) domain.Port {
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
