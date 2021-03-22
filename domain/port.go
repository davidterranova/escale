package domain

import "errors"

var (
	ErrPortNotFound = errors.New("port not found")
)

type Port struct {
	ID       string
	Name     string
	City     string
	Country  string
	Alias    []string
	Regions  []string
	Province string
	Coords   []float32 `json:"coordinates"`
	Timezone string
	Unlocs   []string
	Code     string
}
