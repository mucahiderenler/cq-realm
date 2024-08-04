package handlers

import (
	"encoding/json"
	"mucahiderenler/conquerors-realm/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

type MapHandler struct {
	Service *services.MapService
}

func NewMapHandler(service *services.MapService) *MapHandler {
	return &MapHandler{Service: service}
}

func (m *MapHandler) getMapById(w http.ResponseWriter, r *http.Request) {
	Map, err := m.Service.GetMap()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Map)
}

func (m *MapHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/map", m.getMapById).Methods("GET")
}
