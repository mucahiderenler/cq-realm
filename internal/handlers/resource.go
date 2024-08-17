package handlers

import (
	"encoding/json"
	"mucahiderenler/conquerors-realm/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

type ResourceHandler struct {
	Service *services.ResourceService
}

func NewResourceHandler(service *services.ResourceService) *ResourceHandler {
	return &ResourceHandler{Service: service}
}

func (resourceHandler *ResourceHandler) GetVillageResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	villageID := vars["villageId"]

	resources, err := resourceHandler.Service.GetVillageResources(r.Context(), villageID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)
}

func (resourceHandler *ResourceHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/resources/village/{villageId:[a-zA-Z0-9]+}", resourceHandler.GetVillageResource)
}
