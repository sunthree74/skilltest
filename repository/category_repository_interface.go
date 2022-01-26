package repository

import (
	"context"
	"database/sql"
	"sunthree74/skilltest/model/entity"
)

type CategoryRepositoryInterface interface {
	Save(ctx context.Context, tx *sql.Tx, article entity.Category) entity.Category
	Update(ctx context.Context, tx *sql.Tx, article entity.Category) entity.Category
	Delete(ctx context.Context, tx *sql.Tx, article entity.Category)
	FindById(ctx context.Context, tx *sql.Tx, article int) (entity.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Category
}
