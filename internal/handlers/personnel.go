package handlers

import (
	"encoding/json"
	"garrison/internal/models"
	"garrison/internal/stores"
	"net/http"

	"github.com/google/uuid"
)

type PersonnelHandler struct {
	store stores.PersonnelStore
}

func NewPersonnelHandler(store stores.PersonnelStore) *PersonnelHandler {
	return &PersonnelHandler{store: store}
}

func (h *PersonnelHandler) GetPersonnelByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	p, err := h.store.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (h *PersonnelHandler) GetAllPersonnel(w http.ResponseWriter, r *http.Request) {
	personnel, err := h.store.GetAll(r.Context())
	if err != nil {
		http.Error(w, "error retrieving personnel", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(personnel)
}

func (h *PersonnelHandler) CreatePersonnel(w http.ResponseWriter, r *http.Request) {
	var p models.Personnel
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	created, err := h.store.Create(r.Context(), &p)
	if err != nil {
		http.Error(w, "error creating personnel", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *PersonnelHandler) DeletePersonnel(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.store.Delete(r.Context(), id); err != nil {
		http.Error(w, "error deleting personnel", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PersonnelHandler) UpdatePersonnel(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var p models.Personnel
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	p.ID = id

	updated, err := h.store.Update(r.Context(), &p)
	if err != nil {
		http.Error(w, "error updating personnel", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}
