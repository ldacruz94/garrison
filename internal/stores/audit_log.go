package stores

import (
	"context"

	"garrison/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuditLogStore interface {
	GetAll(ctx context.Context) ([]*models.AuditLog, error)
	GetByEntityID(ctx context.Context, entityID uuid.UUID) ([]*models.AuditLog, error)
	Create(ctx context.Context, l *models.AuditLog) (*models.AuditLog, error)
}

type PgAuditLogStore struct {
	db *pgxpool.Pool
}

func NewAuditLogStore(db *pgxpool.Pool) *PgAuditLogStore {
	return &PgAuditLogStore{db: db}
}

func (s *PgAuditLogStore) GetAll(ctx context.Context) ([]*models.AuditLog, error) {
	query := `
		SELECT id, entity_type, entity_id, actor_id, action, old_value, new_value, occured_at
		FROM audit_logs
		ORDER BY occured_at DESC
	`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*models.AuditLog
	for rows.Next() {
		var l models.AuditLog
		err := rows.Scan(&l.ID, &l.EntityType, &l.EntityID, &l.ActorID, &l.Action, &l.OldValue, &l.NewValue, &l.OccuredAt)
		if err != nil {
			return nil, err
		}
		logs = append(logs, &l)
	}

	return logs, nil
}

func (s *PgAuditLogStore) GetByEntityID(ctx context.Context, entityID uuid.UUID) ([]*models.AuditLog, error) {
	query := `
		SELECT id, entity_type, entity_id, actor_id, action, old_value, new_value, occured_at
		FROM audit_logs
		WHERE entity_id = $1
		ORDER BY occured_at DESC
	`

	rows, err := s.db.Query(ctx, query, entityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*models.AuditLog
	for rows.Next() {
		var l models.AuditLog
		err := rows.Scan(&l.ID, &l.EntityType, &l.EntityID, &l.ActorID, &l.Action, &l.OldValue, &l.NewValue, &l.OccuredAt)
		if err != nil {
			return nil, err
		}
		logs = append(logs, &l)
	}

	return logs, nil
}

func (s *PgAuditLogStore) Create(ctx context.Context, l *models.AuditLog) (*models.AuditLog, error) {
	query := `
		INSERT INTO audit_logs (entity_type, entity_id, actor_id, action, old_value, new_value)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, entity_type, entity_id, actor_id, action, old_value, new_value, occured_at
	`

	var created models.AuditLog
	err := s.db.QueryRow(ctx, query, l.EntityType, l.EntityID, l.ActorID, l.Action, l.OldValue, l.NewValue).Scan(
		&created.ID, &created.EntityType, &created.EntityID, &created.ActorID,
		&created.Action, &created.OldValue, &created.NewValue, &created.OccuredAt,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}
