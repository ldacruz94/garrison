package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID         uuid.UUID  `json:"id"`
	EntityType string     `json:"entity_type"`
	EntityID   uuid.UUID  `json:"entity_id"`
	ActorID    uuid.UUID  `json:"actor_id"`
	Action     string     `json:"action"`
	OldValue   string     `json:"old_value"`
	NewValue   string     `json:"new_value"`
	OccuredAt  *time.Time `json:"occured_at"`
}
