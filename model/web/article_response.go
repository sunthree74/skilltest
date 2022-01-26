package web

import (
	"sunthree74/skilltest/model/entity"
	"time"
)

type ArticleResponse struct {
	Id         int             `json:"id"`
	Title      string          `json:"title"`
	Slug       string          `json:"slug"`
	CategoryId int             `json:"category_id"`
	Content    string          `json:"content"`
	Category   entity.Category `json:"category"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}
