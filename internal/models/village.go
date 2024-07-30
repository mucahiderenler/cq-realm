package models

type Village struct {
	ID       string
	PlayerID string
	Name     string
	X        int
	Y        int
	// Resources       Resources
	// ProductionRates ProductionRates
	// Buildings       []Building
	Soldiers     map[string]int // map of unit type to quantity
	DefenseLevel int
}
