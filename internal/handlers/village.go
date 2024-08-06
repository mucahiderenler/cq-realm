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

func (h *VillageHandler) CreateVillage(w http.ResponseWriter, r *http.Request) {
	// var village models.Village
	// if err := json.NewDecoder(r.Body).Decode(&village); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// if err := h.Service.CreateVillage(village); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(village)
}

func (h *VillageHandler) UpdateVillage(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// villageID := vars["id"]

	// var village models.Village
	// if err := json.NewDecoder(r.Body).Decode(&village); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// if err := h.Service.UpdateVillage(&village); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(village)
}

func (h *VillageHandler) DeleteVillage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	villageID := vars["id"]

	if err := h.Service.DeleteVillage(villageID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (v *VillageHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/villages/{id:[a-zA-Z0-9]+}", v.GetVillage).Methods("GET")
	r.HandleFunc("/villages", v.CreateVillage).Methods("POST")
	r.HandleFunc("/villages/{id:[a-zA-Z0-9]+}", v.UpdateVillage).Methods("PUT")
	r.HandleFunc("/villages/{id:[a-zA-Z0-9]+}", v.DeleteVillage).Methods("DELETE")
}
