package entity

import "time"

type Category struct {
	Id           int
	CategoryName string
	CategorySlug string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
