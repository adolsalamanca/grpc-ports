package entity

import (
	any "github.com/golang/protobuf/ptypes/any"
)

type Port map[string]PortInfo

type PortInfo struct {
	Name        string        `json:"name"`
	City        string        `json:"city"`
	Country     string        `json:"country"`
	Alias       []*any.Any      `json:"alias"`
	Regions     []*any.Any `json:"regions"`
	Coordinates []float32     `json:"coordinates"`
	Province    *string       `json:"province,omitempty"`
	Timezone    *string       `json:"timezone,omitempty"`
	Unlocs      []string      `json:"unlocs"`
	Code        *string       `json:"code,omitempty"`
}