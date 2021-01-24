package entity


type Port map[string]PortInfo

type PortInfo struct {
	Name        string        `json:"name"`
	City        string        `json:"city"`
	Country     string        `json:"country"`
	Alias       []string      `json:"alias"`
	Regions     []string `json:"regions"`
	Coordinates []float64     `json:"coordinates"`
	Province    *string       `json:"province,omitempty"`
	Timezone    *string       `json:"timezone,omitempty"`
	Unlocs      []string      `json:"unlocs"`
	Code        *string       `json:"code,omitempty"`
}