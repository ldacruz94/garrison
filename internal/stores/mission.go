package stores

import (
	"context"
	"errors"
	"garrison/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionStore interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Mission, error)
	GetAll(ctx context.Context) ([]*models.Mission, error)
	Create(ctx context.Context, m *models.Mission) (*models.Mission, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, m *models.Mission) (*models.Mission, error)
}

type PgMissionStore struct {
	db *pgxpool.Pool
}

func NewMissionStore(db *pgxpool.Pool) *PgMissionStore {
	return &PgMissionStore{db: db}
}

func (s *PgMissionStore) GetByID(ctx context.Context, id uuid.UUID) (*models.Mission, error) {
	query := `
		SELECT id, name, description, status, mission_type, start_time, end_time, created_at, updated_at
		FROM missions
		WHERE id = $1
	`

	var m models.Mission
	err := s.db.QueryRow(ctx, query, id).Scan(
		&m.ID, &m.Name, &m.Description, &m.Status, &m.MissionType,
		&m.StartTime, &m.EndTime, &m.CreatedAt, &m.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("mission not found")
		}
		return nil, err
	}

	return &m, nil
}

func (s *PgMissionStore) GetAll(ctx context.Context) ([]*models.Mission, error) {
	query := `
		SELECT id, name, description, status, mission_type, start_time, end_time, created_at, updated_at
		FROM missions
	`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var missions []*models.Mission
	for rows.Next() {
		var m models.Mission
		err := rows.Scan(
			&m.ID, &m.Name, &m.Description, &m.Status, &m.MissionType,
			&m.StartTime, &m.EndTime, &m.CreatedAt, &m.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		missions = append(missions, &m)
	}

	return missions, nil
}

func (s *PgMissionStore) Create(ctx context.Context, m *models.Mission) (*models.Mission, error) {
	query := `
		INSERT INTO missions (name, description, status, mission_type, start_time, end_time)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, description, status, mission_type, start_time, end_time, created_at, updated_at
	`

	var created models.Mission
	err := s.db.QueryRow(ctx, query, m.Name, m.Description, m.Status, m.MissionType, m.StartTime, m.EndTime).Scan(
		&created.ID, &created.Name, &created.Description, &created.Status, &created.MissionType,
		&created.StartTime, &created.EndTime, &created.CreatedAt, &created.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (s *PgMissionStore) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM missions WHERE id = $1`

	_, err := s.db.Exec(ctx, query, id)
	return err
}

func (s *PgMissionStore) Update(ctx context.Context, m *models.Mission) (*models.Mission, error) {
	query := `
		UPDATE missions
		SET name = $1, description = $2, status = $3, mission_type = $4, start_time = $5, end_time = $6, updated_at = NOW()
		WHERE id = $7
		RETURNING id, name, description, status, mission_type, start_time, end_time, created_at, updated_at
	`

	var updated models.Mission
	err := s.db.QueryRow(ctx, query, m.Name, m.Description, m.Status, m.MissionType, m.StartTime, m.EndTime, m.ID).Scan(
		&updated.ID, &updated.Name, &updated.Description, &updated.Status, &updated.MissionType,
		&updated.StartTime, &updated.EndTime, &updated.CreatedAt, &updated.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("mission not found")
		}
		return nil, err
	}

	return &updated, nil
}
