package domain

type Port struct {
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
