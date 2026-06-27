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
		SELECT id, name, status, created_at 
		FROM missions 
		WHERE id = $1
	`

	var m models.Mission
	row := s.db.QueryRow(ctx, query, id)
	err := row.Scan(&m.Name, &m.Status, &m.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *PgMissionStore) GetAll(ctx context.Context) ([]*models.Mission, error) {
	query := `SELECT id, name, status, created_at FROM missions`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var missions []*models.Mission

	for rows.Next() {
		var m models.Mission
		err := rows.Scan(&m.ID, &m.Name, &m.Status, &m.CreatedAt)
		if err != nil {
			return nil, err
		}

		missions = append(missions, &m)
	}

	return missions, nil
}

func (s *PgMissionStore) Create(ctx context.Context, m *models.Mission) (*models.Mission, error) {
	query := `
		INSERT INTO missions (name, status, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, name, status, created_at
	`

	var created models.Mission
	err := s.db.QueryRow(ctx, query, m.Name, m.Status, m.CreatedAt).Scan(
		&created.ID,
		&created.Name,
		&created.Status,
		&created.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (s *PgMissionStore) Delete(ctx context.Context, id uuid.UUID) error {

	query := `DELETE From missions WHERE id = $1`

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PgMissionStore) Update(ctx context.Context, m *models.Mission) (*models.Mission, error) {
	query := `
		UPDATE missions
		SET name = $1, status = $2
		WHERE id = $3
		RETURNING id, name, status, created_at
	`

	var updated models.Mission
	err := s.db.QueryRow(ctx, query, m.Name, m.Status, m.ID).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Status,
		&updated.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("mission not found")
		}
		return nil, err
	}

	return &updated, nil
}
