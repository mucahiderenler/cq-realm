package types

const (
	TypeBuildingUpgrade = "building:upgrade"
	TyepUnitRecruit     = "unit:recruit"
)

type BuildingUpgradePayload struct {
	VillageID  string
	BuildingID string
}
