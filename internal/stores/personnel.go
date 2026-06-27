package stores

import (
	"context"
	"errors"

	"garrison/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PersonnelStore interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Personnel, error)
	GetAll(ctx context.Context) ([]*models.Personnel, error)
	Create(ctx context.Context, p *models.Personnel) (*models.Personnel, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, p *models.Personnel) (*models.Personnel, error)
}

type PgPersonnelStore struct {
	db *pgxpool.Pool
}

func NewPersonnelStore(db *pgxpool.Pool) *PgPersonnelStore {
	return &PgPersonnelStore{db: db}
}

func (s *PgPersonnelStore) GetByID(ctx context.Context, id uuid.UUID) (*models.Personnel, error) {
	query := `
		SELECT id, rank, last_name, first_name, unit_designator, clearance_level, status, created_at, updated_at
		FROM personnel
		WHERE id = $1
	`

	var p models.Personnel
	err := s.db.QueryRow(ctx, query, id).Scan(
		&p.ID, &p.Rank, &p.LastName, &p.FirstName,
		&p.UnitDesignator, &p.ClearanceLevel, &p.Status,
		&p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("personnel not found")
		}
		return nil, err
	}

	return &p, nil
}

func (s *PgPersonnelStore) GetAll(ctx context.Context) ([]*models.Personnel, error) {
	query := `
		SELECT id, rank, last_name, first_name, unit_designator, clearance_level, status, created_at, updated_at
		FROM personnel
	`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var personnel []*models.Personnel
	for rows.Next() {
		var p models.Personnel
		err := rows.Scan(
			&p.ID, &p.Rank, &p.LastName, &p.FirstName,
			&p.UnitDesignator, &p.ClearanceLevel, &p.Status,
			&p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		personnel = append(personnel, &p)
	}

	return personnel, nil
}

func (s *PgPersonnelStore) Create(ctx context.Context, p *models.Personnel) (*models.Personnel, error) {
	query := `
		INSERT INTO personnel (rank, last_name, first_name, unit_designator, clearance_level, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, rank, last_name, first_name, unit_designator, clearance_level, status, created_at, updated_at
	`

	var created models.Personnel
	err := s.db.QueryRow(ctx, query, p.Rank, p.LastName, p.FirstName, p.UnitDesignator, p.ClearanceLevel, p.Status).Scan(
		&created.ID, &created.Rank, &created.LastName, &created.FirstName,
		&created.UnitDesignator, &created.ClearanceLevel, &created.Status,
		&created.CreatedAt, &created.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (s *PgPersonnelStore) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM personnel WHERE id = $1`

	result, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("personnel not found")
	}

	return nil
}

func (s *PgPersonnelStore) Update(ctx context.Context, p *models.Personnel) (*models.Personnel, error) {
	query := `
		UPDATE personnel
		SET rank = $1, last_name = $2, first_name = $3, unit_designator = $4, clearance_level = $5, status = $6
		WHERE id = $7
		RETURNING id, rank, last_name, first_name, unit_designator, clearance_level, status, created_at, updated_at
	`

	var updated models.Personnel
	err := s.db.QueryRow(ctx, query, p.Rank, p.LastName, p.FirstName, p.UnitDesignator, p.ClearanceLevel, p.Status, p.ID).Scan(
		&updated.ID, &updated.Rank, &updated.LastName, &updated.FirstName,
		&updated.UnitDesignator, &updated.ClearanceLevel, &updated.Status,
		&updated.CreatedAt, &updated.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("personnel not found")
		}
		return nil, err
	}

	return &updated, nil
}
