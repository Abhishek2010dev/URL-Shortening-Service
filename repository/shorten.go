package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Abhishek2010dev/URL-Shortening-Service/dto"
	"github.com/Abhishek2010dev/URL-Shortening-Service/model"
	"github.com/jmoiron/sqlx"
)

var ErrShortenNotFound = errors.New("Shorten not found")

type CreateShortenPayload struct {
	Url       string `db:"url"`
	ShortCode string `db:"short_code"`
}

type Shorten interface {
	Create(ctx context.Context, payload CreateShortenPayload) (*model.Shorten, error)
	FindByShortCode(ctx context.Context, shortCode string) (*model.Shorten, error)
	FindByShortCodeWithAccessCount(ctx context.Context, shortCode string) (*model.ShortenWithAccessCount, error)
	Delete(ctx context.Context, shortCode string) error
	Update(ctx context.Context, shortCode string, payload dto.ShortenPayload) (*model.Shorten, error)
	IncrementAccessCount(ctx context.Context, shortCode string) error
}

type shortenRepo struct {
	db *sqlx.DB
}

func NewShorten(db *sqlx.DB) Shorten {
	return &shortenRepo{db}
}

func (shortenrepo *shortenRepo) Create(ctx context.Context, payload CreateShortenPayload) (*model.Shorten, error) {
	query := `
		INSERT INTO shorten (url, short_code) 
	        VALUES (:url, :short_code)
		RETURNING id, url, short_code, created_at, updated_at
	`
	stmt, err := shortenrepo.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	var shorten model.Shorten
	if err = stmt.GetContext(ctx, &shorten, payload); err != nil {
		return nil, fmt.Errorf("failed to create shorten url: %w", err)
	}

	return &shorten, nil
}

func (shortenrepo *shortenRepo) FindByShortCode(ctx context.Context, shortCode string) (*model.Shorten, error) {
	query := `
		SELECT id, url, short_code, created_at, updated_at FROM shorten 
		WHERE short_code = $1
	`

	stmt, err := shortenrepo.db.PreparexContext(ctx, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrShortenNotFound
		}
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

	var shorten model.Shorten
	if err := stmt.GetContext(ctx, &shorten, shortCode); err != nil {
		return nil, fmt.Errorf("failed to get shorten (short_code: %v): %w", shortCode, err)
	}

	return &shorten, nil
}

func (shortenrepo *shortenRepo) FindByShortCodeWithAccessCount(ctx context.Context, shortCode string) (*model.ShortenWithAccessCount, error) {
	query := `
		SELECT id, url, short_code, access_count, created_at, updated_at FROM shorten 
		WHERE short_code = $1
	`

	stmt, err := shortenrepo.db.PreparexContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

	var shortenWithAccessCount model.ShortenWithAccessCount
	if err := stmt.GetContext(ctx, &shortenWithAccessCount, shortCode); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrShortenNotFound
		}
		return nil, fmt.Errorf("failed to get shorten (short_code: %v): %w", shortCode, err)
	}

	return &shortenWithAccessCount, nil

}
func (shortenrepo *shortenRepo) Delete(ctx context.Context, shortCode string) error {
	panic("not implemented") // TODO: Implement
}
func (shortenrepo *shortenRepo) Update(ctx context.Context, shortCode string, payload dto.ShortenPayload) (*model.Shorten, error) {
	panic("not implemented") // TODO: Implement
}
func (shortenrepo *shortenRepo) IncrementAccessCount(ctx context.Context, shortCode string) error {
	panic("not implemented") // TODO: Implement
}
