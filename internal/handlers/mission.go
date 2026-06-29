package handlers

import (
	"encoding/json"
	"garrison/internal/models"
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

func (h *MissionHandler) GetAllMissions(w http.ResponseWriter, r *http.Request) {
	missions, err := h.store.GetAll(r.Context())

	if err != nil {
		http.Error(w, "error retrieving missions", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(missions)
}

func (h *MissionHandler) CreateMission(w http.ResponseWriter, r *http.Request) {
	var m models.Mission
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	newMission, err := h.store.Create(r.Context(), &m)
	if err != nil {
		http.Error(w, "error in creating new mission", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMission)
}

func (h *MissionHandler) DeleteMission(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "id cannot be empty", http.StatusBadRequest)
	}

	parsed_id, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.store.Delete(r.Context(), parsed_id)
	if err != nil {
		http.Error(w, "error deleting mission", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *MissionHandler) UpdateMission(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var m models.Mission
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	m.ID = id

	updatedMission, err := h.store.Update(r.Context(), &m)
	if err != nil {
		http.Error(w, "error in creating new mission", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedMission)
}
