package web

import "time"

type CategoryResponse struct {
	Id           int       `json:"id"`
	CategoryName string    `json:"category_name"`
	CategorySlug string    `json:"category_slug"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
