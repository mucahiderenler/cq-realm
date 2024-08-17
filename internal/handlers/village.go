package handlers

import (
	"encoding/json"
	"mucahiderenler/conquerors-realm/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

type VillageHandler struct {
	Service *services.VillageService
}

func NewVillageHandler(service *services.VillageService) *VillageHandler {
	return &VillageHandler{Service: service}
}

func (h *VillageHandler) GetVillage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	villageID := vars["id"]

	villageResult, err := h.Service.GetVillageByID(r.Context(), villageID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(villageResult)
}

func (v *VillageHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/villages/{id:[a-zA-Z0-9]+}", v.GetVillage).Methods("GET")
}
