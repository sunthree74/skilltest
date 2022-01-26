package entity

import "time"

type Article struct {
	Id         int
	Title      string
	Slug       string
	CategoryId int
	Category   Category
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
