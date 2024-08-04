package models

type VillageID string
type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Map struct {
	Villages []*Village `json:"villages"`
}
