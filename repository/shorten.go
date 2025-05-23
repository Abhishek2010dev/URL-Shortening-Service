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
	const query = `
        INSERT INTO shorten (url, short_code) 
        VALUES (:url, :short_code)
        RETURNING id, url, short_code, created_at, updated_at
    `
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("prepare insert statement: %w", err)
	}
	defer stmt.Close()

	var shorten model.Shorten
	if err := stmt.GetContext(ctx, &shorten, payload); err != nil {
		return nil, fmt.Errorf("execute insert: %w", err)
	}
	return &shorten, nil
}

func (r *shortenRepo) FindByShortCode(ctx context.Context, shortCode string) (*model.Shorten, error) {
	const query = `
        SELECT id, url, short_code, created_at, updated_at 
        FROM shorten 
        WHERE short_code = $1
    `
	var shorten model.Shorten
	if err := r.db.GetContext(ctx, &shorten, query, shortCode); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrShortCodeNotFound
		}
		return nil, fmt.Errorf("get shorten by short_code: %w", err)
	}
	return &shorten, nil
}

func (r *shortenRepo) FindByShortCodeWithAccessCount(ctx context.Context, shortCode string) (*model.ShortenWithAccessCount, error) {
	const query = `
        SELECT id, url, short_code, access_count, created_at, updated_at 
        FROM shorten 
        WHERE short_code = $1
    `
	var shorten model.ShortenWithAccessCount
	if err := r.db.GetContext(ctx, &shorten, query, shortCode); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrShortCodeNotFound
		}
		return nil, fmt.Errorf("get shorten with access count: %w", err)
	}
	return &shorten, nil
}

func (r *shortenRepo) Delete(ctx context.Context, shortCode string) error {
	const query = `DELETE FROM shorten WHERE short_code = $1`
	result, err := r.db.ExecContext(ctx, query, shortCode)
	if err != nil {
		return fmt.Errorf("delete shorten: %w", err)
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return ErrShortCodeNotFound
	}
	return nil
}

func (r *shortenRepo) Update(ctx context.Context, payload ShortenPayload) (*model.Shorten, error) {
	const query = `
        UPDATE shorten 
        SET url = :url 
        WHERE short_code = :short_code
        RETURNING id, url, short_code, created_at, updated_at
    `
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("prepare update statement: %w", err)
	}
	defer stmt.Close()

	var shorten model.Shorten
	if err := stmt.GetContext(ctx, &shorten, payload); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrShortCodeNotFound
		}
		return nil, fmt.Errorf("execute update: %w", err)
	}
	return &shorten, nil
}

func (r *shortenRepo) IncrementAccessCount(ctx context.Context, shortCode string) error {
	const query = `UPDATE shorten SET access_count = access_count + 1 WHERE short_code = $1`
	result, err := r.db.ExecContext(ctx, query, shortCode)
	if err != nil {
		return fmt.Errorf("increment access count: %w", err)
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return ErrShortCodeNotFound
	}
	return nil
}
