package handlers

import (
	"encoding/json"
	"garrison/internal/models"
	"garrison/internal/stores"
	"net/http"

	"github.com/google/uuid"
)

type AssetHandler struct {
	store stores.AssetStore
}

func NewAssetHandler(store stores.AssetStore) *AssetHandler {
	return &AssetHandler{store: store}
}

func (h *AssetHandler) GetAssetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	a, err := h.store.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

func (h *AssetHandler) GetAllAssets(w http.ResponseWriter, r *http.Request) {
	assets, err := h.store.GetAll(r.Context())
	if err != nil {
		http.Error(w, "error retrieving assets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assets)
}

func (h *AssetHandler) CreateAsset(w http.ResponseWriter, r *http.Request) {
	var a models.Asset
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	created, err := h.store.Create(r.Context(), &a)
	if err != nil {
		http.Error(w, "error creating asset", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *AssetHandler) DeleteAsset(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.store.Delete(r.Context(), id); err != nil {
		http.Error(w, "error deleting asset", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AssetHandler) UpdateAsset(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var a models.Asset
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	a.ID = id

	updated, err := h.store.Update(r.Context(), &a)
	if err != nil {
		http.Error(w, "error updating asset", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}
