package model

import "time"

type Shorten struct {
	Id        int       `json:"id" db:"id"`
	Url       string    `json:"url" db:"url"`
	ShortCode string    `json:"short_code" db:"short_code"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ShortenWithAccessCount struct {
	Shorten
	AccessCount int `json:"access_count" db:"access_count"`
}
