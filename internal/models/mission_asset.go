package models

import (
	"time"

	"github.com/google/uuid"
)

type MissionAsset struct {
	ID         uuid.UUID  `json:"id"`
	MissionID  uuid.UUID  `json:"mission_id"`
	AssetID    uuid.UUID  `json:"asset_id"`
	Role       string     `json:"role"`
	AssignedAt *time.Time `json:"AssignedAt"`
}
