package repository

import (
	"context"

	"github.com/Abhishek2010dev/URL-Shortening-Service/dto"
	"github.com/Abhishek2010dev/URL-Shortening-Service/model"
)

type Shorten interface {
	Create(ctx context.Context, payload dto.ShortenPayload) (*model.Shorten, error)
	FindByShortCode(ctx context.Context, shortCode string) (*model.Shorten, error)
	Delete(ctx context.Context, shortCode string) error
	Update(ctx context.Context, shortCode string, payload dto.ShortenPayload) (*model.Shorten, error)
	FindAll(ctx context.Context) ([]model.Shorten, error)
	IncrementAccessCount(ctx context.Context, shortCode string) error
}
