package models

type VillageID string
type Coordinates struct {
	X int
	Y int
}
type Map struct {
	ID       int
	Name     string
	Villages map[VillageID]Coordinates
}
