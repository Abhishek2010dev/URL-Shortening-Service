package handler

import (
	"errors"

	"github.com/Abhishek2010dev/URL-Shortening-Service/repository"
	"github.com/Abhishek2010dev/URL-Shortening-Service/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Shorten struct {
	repo repository.Shorten
}

func NewShorten(repo repository.Shorten) *Shorten {
	return &Shorten{repo}
}

type ShortenPayload struct {
	Url string `json:"url"`
}

func (s *Shorten) Create(c fiber.Ctx) error {
	var payload ShortenPayload
	if err := c.Bind().JSON(&payload); err != nil {
		return err
	}
	if !utils.IsValidateUrl(payload.Url) {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid url format")
	}
	createPaylaod := repository.ShortenPayload{
		ShortCode: uuid.New().String(),
		URL:       payload.Url,
	}
	shorten, err := s.repo.Create(c.Context(), createPaylaod)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(shorten)
}

func (s *Shorten) GetByShortCode(c fiber.Ctx) error {
	shortCode := c.Params("short_code")
	responseBody, err := s.repo.FindByShortCode(c.Context(), shortCode)
	if err != nil {
		if errors.Is(err, repository.ErrShortCodeNotFound) {
			return fiber.ErrNotFound
		}
		return err
	}
	return c.JSON(responseBody)
}

func (s *Shorten) GetURLStatistics(c fiber.Ctx) error {
	shortCode := c.Params("short_code")
	responseBody, err := s.repo.FindByShortCodeWithAccessCount(c.Context(), shortCode)
	if err != nil {
		if errors.Is(err, repository.ErrShortCodeNotFound) {
			return fiber.ErrNotFound
		}
		return err
	}
	return c.JSON(responseBody)
}

func (s *Shorten) Delete(c fiber.Ctx) error {
	shortCode := c.Params("short_code")
	if err := s.repo.Delete(c.Context(), shortCode); err != nil {
		if errors.Is(err, repository.ErrShortCodeNotFound) {
			return fiber.ErrNotFound
		}
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
