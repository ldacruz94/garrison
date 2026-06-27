package handlers

import (
	"encoding/json"
	"garrison/internal/stores"
	"net/http"

	"github.com/google/uuid"
)

type MissionHandler struct {
	store stores.MissionStore
}

func NewMissionHandler(store stores.MissionStore) *MissionHandler {
	return &MissionHandler{
		store: store,
	}
}

func (h *MissionHandler) GetMissionByID(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "id cannot be empty", http.StatusBadRequest)
	}

	parsed_id, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	mission, err := h.store.GetByID(r.Context(), parsed_id)

	if err != nil {
		http.Error(w, "Mission retrieval failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mission)
}
