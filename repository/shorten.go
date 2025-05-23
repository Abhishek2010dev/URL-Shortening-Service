package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Abhishek2010dev/URL-Shortening-Service/model"
	"github.com/jmoiron/sqlx"
)

var ErrShortCodeNotFound = errors.New("short_code not found")

type ShortenPayload struct {
	URL       string `db:"url"`
	ShortCode string `db:"short_code"`
}

type Shorten interface {
	Create(ctx context.Context, payload ShortenPayload) (*model.Shorten, error)
	FindByShortCode(ctx context.Context, shortCode string) (*model.Shorten, error)
	FindByShortCodeWithAccessCount(ctx context.Context, shortCode string) (*model.ShortenWithAccessCount, error)
	Delete(ctx context.Context, shortCode string) error
	Update(ctx context.Context, payload ShortenPayload) (*model.Shorten, error)
	IncrementAccessCount(ctx context.Context, shortCode string) error
}

type shortenRepo struct {
	db *sqlx.DB
}

func NewShorten(db *sqlx.DB) Shorten {
	return &shortenRepo{db}
}

func (r *shortenRepo) Create(ctx context.Context, payload ShortenPayload) (*model.Shorten, error) {
	query := `
        INSERT INTO shorten (url, short_code) 
        VALUES (:url, :short_code)
        RETURNING id, url, short_code, created_at, updated_at
    `
	rows, err := r.db.NamedQueryContext(ctx, query, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to execute insert query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var shorten model.Shorten
		if err := rows.StructScan(&shorten); err != nil {
			return nil, fmt.Errorf("failed to scan inserted row: %w", err)
		}
		return &shorten, nil
	}

	return nil, fmt.Errorf("no row returned after insert")
}

func (r *shortenRepo) FindByShortCode(ctx context.Context, shortCode string) (*model.Shorten, error) {
	query := `
        SELECT id, url, short_code, created_at, updated_at 
        FROM shorten 
        WHERE short_code = $1
    `
	var shorten model.Shorten
	err := r.db.GetContext(ctx, &shorten, query, shortCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrShortCodeNotFound
		}
		return nil, fmt.Errorf("failed to get shorten (short_code: %v): %w", shortCode, err)
	}

	return &shorten, nil
}

func (r *shortenRepo) FindByShortCodeWithAccessCount(ctx context.Context, shortCode string) (*model.ShortenWithAccessCount, error) {
	query := `
        SELECT id, url, short_code, access_count, created_at, updated_at 
        FROM shorten 
        WHERE short_code = $1
    `
	var shorten model.ShortenWithAccessCount
	err := r.db.GetContext(ctx, &shorten, query, shortCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrShortCodeNotFound
		}
		return nil, fmt.Errorf("failed to get shorten with access count (short_code: %v): %w", shortCode, err)
	}

	return &shorten, nil
}

func (r *shortenRepo) Delete(ctx context.Context, shortCode string) error {
	query := "DELETE FROM shorten WHERE short_code = $1"
	result, err := r.db.ExecContext(ctx, query, shortCode)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return ErrShortCodeNotFound
	}

	return nil
}

func (r *shortenRepo) Update(ctx context.Context, payload ShortenPayload) (*model.Shorten, error) {
	query := `
        UPDATE shorten 
        SET url = :url 
        WHERE short_code = :short_code
        RETURNING id, url, short_code, created_at, updated_at
    `
	rows, err := r.db.NamedQueryContext(ctx, query, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to execute update query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var shorten model.Shorten
		if err := rows.StructScan(&shorten); err != nil {
			return nil, fmt.Errorf("failed to scan updated row: %w", err)
		}
		return &shorten, nil
	}

	return nil, ErrShortCodeNotFound
}

func (r *shortenRepo) IncrementAccessCount(ctx context.Context, shortCode string) error {
	query := "UPDATE shorten SET access_count = access_count + 1 WHERE short_code = $1"
	result, err := r.db.ExecContext(ctx, query, shortCode)
	if err != nil {
		return fmt.Errorf("failed to execute increment query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return ErrShortCodeNotFound
	}

	return nil
}
