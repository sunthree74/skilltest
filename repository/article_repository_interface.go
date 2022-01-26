package repository

import (
	"context"
	"database/sql"
	"sunthree74/skilltest/model/entity"
)

type ArticleRepositoryInterface interface {
	Save(ctx context.Context, tx *sql.Tx, article entity.Article) entity.Article
	Update(ctx context.Context, tx *sql.Tx, article entity.Article) entity.Article
	Delete(ctx context.Context, tx *sql.Tx, article entity.Article)
	FindById(ctx context.Context, tx *sql.Tx, article int) (entity.Article, error)
	FindAll(ctx context.Context, tx *sql.Tx, QueryTitle string, QueryCategoryName string) []entity.Article
}
