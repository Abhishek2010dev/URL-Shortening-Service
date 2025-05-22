package model

import "time"

type Shorten struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	ShortCode string    `json:"short_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ShortenWithAccessCount struct {
	Shorten
	AccessCount int `json:"access_count"`
}
