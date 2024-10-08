package handlers

import (
	"encoding/json"
	"io"
	"mucahiderenler/conquerors-realm/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

type UpgradeBuildingBody struct {
	VillageId string `json:"villageId"`
}

type BuildingHandler struct {
	buildingService *services.BuildingService
}

func NewBuildingHandler(buildingService *services.BuildingService) *BuildingHandler {
	return &BuildingHandler{buildingService: buildingService}
}

func (h *BuildingHandler) UpgradeBuilding(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	buildingId := vars["buildingId"]

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var requestBody UpgradeBuildingBody

	err = json.Unmarshal(body, &requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.buildingService.UpgradeBuildingInit(r.Context(), buildingId, requestBody.VillageId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *BuildingHandler) GetBuildingDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	buildingId := vars["buildingId"]
	villageId := vars["villageId"]

	buildingConfig, err := h.buildingService.GetBuildingDetails(r.Context(), buildingId, villageId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buildingConfig)
}

func (h *BuildingHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/building/{buildingId:[a-zA-Z0-9]+}/upgrade", h.UpgradeBuilding).Methods("POST")
	r.HandleFunc("/villages/{villageId:[a-zA-Z0-9]+}/building/{buildingId:[a-zA-Z0-9]+}", h.GetBuildingDetails).Methods("GET")
}
