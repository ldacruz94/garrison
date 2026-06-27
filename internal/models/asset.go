package models

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	ID          uuid.UUID  `json:"id"`
	Designation string     `json:"designation"`
	AssetType   string     `json:"asset_type"`
	Notes       string     `json:"notes"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
