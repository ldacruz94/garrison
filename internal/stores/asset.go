package stores

import (
	"context"
	"errors"

	"garrison/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AssetStore interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Asset, error)
	GetAll(ctx context.Context) ([]*models.Asset, error)
	Create(ctx context.Context, a *models.Asset) (*models.Asset, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, a *models.Asset) (*models.Asset, error)
}

type PgAssetStore struct {
	db *pgxpool.Pool
}

func NewAssetStore(db *pgxpool.Pool) *PgAssetStore {
	return &PgAssetStore{db: db}
}

func (s *PgAssetStore) GetByID(ctx context.Context, id uuid.UUID) (*models.Asset, error) {
	query := `
		SELECT id, designation, asset_type, notes, created_at, updated_at
		FROM assets
		WHERE id = $1
	`

	var a models.Asset
	err := s.db.QueryRow(ctx, query, id).Scan(
		&a.ID, &a.Designation, &a.AssetType, &a.Notes, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("asset not found")
		}
		return nil, err
	}

	return &a, nil
}

func (s *PgAssetStore) GetAll(ctx context.Context) ([]*models.Asset, error) {
	query := `SELECT id, designation, asset_type, notes, created_at, updated_at FROM assets`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []*models.Asset
	for rows.Next() {
		var a models.Asset
		err := rows.Scan(&a.ID, &a.Designation, &a.AssetType, &a.Notes, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &a)
	}

	return assets, nil
}

func (s *PgAssetStore) Create(ctx context.Context, a *models.Asset) (*models.Asset, error) {
	query := `
		INSERT INTO assets (designation, asset_type, notes)
		VALUES ($1, $2, $3)
		RETURNING id, designation, asset_type, notes, created_at, updated_at
	`

	var created models.Asset
	err := s.db.QueryRow(ctx, query, a.Designation, a.AssetType, a.Notes).Scan(
		&created.ID, &created.Designation, &created.AssetType, &created.Notes,
		&created.CreatedAt, &created.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (s *PgAssetStore) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM assets WHERE id = $1`

	result, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("asset not found")
	}

	return nil
}

func (s *PgAssetStore) Update(ctx context.Context, a *models.Asset) (*models.Asset, error) {
	query := `
		UPDATE assets
		SET designation = $1, asset_type = $2, notes = $3
		WHERE id = $4
		RETURNING id, designation, asset_type, notes, created_at, updated_at
	`

	var updated models.Asset
	err := s.db.QueryRow(ctx, query, a.Designation, a.AssetType, a.Notes, a.ID).Scan(
		&updated.ID, &updated.Designation, &updated.AssetType, &updated.Notes,
		&updated.CreatedAt, &updated.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("asset not found")
		}
		return nil, err
	}

	return &updated, nil
}
