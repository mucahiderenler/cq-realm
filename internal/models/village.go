package models

type VillageType int

// Declare related constants for each weekday starting with index 1
const (
	Normal          VillageType = iota + 1 // EnumIndex = 1
	IronBonus                              // EnumIndex = 2
	WoodBonus                              // EnumIndex = 3
	ClayBonus                              // EnumIndex = 4
	PopulationBonus                        // EnumIndex = 5
	StableBonus                            // EnumIndex = 6
	BarrackBonus                           // EnumIndex = 7
)

type Village struct {
	ID         string      `json:"id"`
	PlayerID   string      `json:"playerID"`
	PlayerName string      `json:"playerName"`
	Name       string      `json:"name"`
	X          int         `json:"x"`
	Y          int         `json:"y"`
	Point      int         `json:"point"`
	Type       VillageType `json:"type"`
	Owner_name string      `json:"owner_name"`
	Owner_id   string      `json:"owner_id"`
	// Resources       Resources
	// ProductionRates ProductionRates
	// Buildings       []Building
	// Soldiers     map[string]int // map of unit type to quantity
	// DefenseLevel int
}
