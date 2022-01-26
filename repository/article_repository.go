package repository

import (
	"context"
	"database/sql"
	"errors"
	"sunthree74/skilltest/helper"
	"sunthree74/skilltest/model/entity"
	"time"
)

type ArticleRepository struct {
}

func NewArticleRepositoryInterface() ArticleRepositoryInterface {
	return &ArticleRepository{}
}

func (repository *ArticleRepository) Save(ctx context.Context, tx *sql.Tx, article entity.Article) entity.Article {
	SQL := "insert into article(title, slug, category_id, content, created_at, updated_at) values (?,?,?,?,?,?)"
	now := time.Now()
	now = now.Local()
	dateNow := now.Format("2006-01-02 15:04:05")
	result, err := tx.ExecContext(ctx, SQL, article.Title, article.Slug, article.CategoryId, article.Content, dateNow, dateNow)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	article.Id = int(id)
	return article
}

func (repository *ArticleRepository) Update(ctx context.Context, tx *sql.Tx, article entity.Article) entity.Article {
	SQL := "update article set title = ?, slug = ?, category_id = ?, content = ?, updated_at = ?  where id = ? and deleted_at is null"
	now := time.Now()
	now = now.Local()
	dateNow := now.Format("2006-01-02 15:04:05")
	_, err := tx.ExecContext(ctx, SQL, article.Title, article.Slug, article.CategoryId,
		article.Content, dateNow, article.Id)
	helper.PanicIfError(err)

	return article
}

func (repository *ArticleRepository) Delete(ctx context.Context, tx *sql.Tx, article entity.Article) {
	SQL := "update article set deleted_at = ?  where id = ? and deleted_at is null"
	now := time.Now()
	now = now.Local()
	dateNow := now.Format("2006-01-02 15:04:05")
	_, err := tx.ExecContext(ctx, SQL, dateNow, article.Id)
	helper.PanicIfError(err)
}

func (repository *ArticleRepository) FindById(ctx context.Context, tx *sql.Tx, articleId int) (entity.Article, error) {
	SQL := "select a.id, a.title, a.slug, a.category_id, a.content, a.created_at, a.updated_at, b.id, b.category_name, b.category_slug, b.created_at, b.updated_at from article a, category b where a.category_id = b.id and a.id = ? and a.deleted_at is null"
	rows, err := tx.QueryContext(ctx, SQL, articleId)
	helper.PanicIfError(err)
	defer rows.Close()

	article := entity.Article{}
	if rows.Next() {
		err := rows.Scan(&article.Id, &article.Title, &article.Slug, &article.CategoryId,
			&article.Content, &article.CreatedAt, &article.UpdatedAt, &article.Category.Id,
			&article.Category.CategoryName, &article.Category.CategorySlug, &article.Category.CreatedAt, &article.Category.UpdatedAt)
		helper.PanicIfError(err)
		return article, nil
	} else {
		return article, errors.New("article is not found")
	}
}

func (repository *ArticleRepository) FindAll(ctx context.Context, tx *sql.Tx, QueryTitle string, QueryCategoryName string) []entity.Article {
	SQL := "select a.id, a.title, a.slug, a.category_id, a.content, a.created_at, a.updated_at, b.id, b.category_name, b.category_slug, b.created_at, b.updated_at from article a, category b where a.category_id = b.id and (a.title like '%" + QueryTitle + "%' or b.category_name like '%" + QueryCategoryName + "%') and a.deleted_at is null group by a.id order by a.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var articles []entity.Article
	for rows.Next() {
		article := entity.Article{}
		err := rows.Scan(&article.Id, &article.Title, &article.Slug, &article.CategoryId,
			&article.Content, &article.CreatedAt, &article.UpdatedAt, &article.Category.Id,
			&article.Category.CategoryName, &article.Category.CategorySlug, &article.Category.CreatedAt, &article.Category.UpdatedAt)
		helper.PanicIfError(err)
		articles = append(articles, article)
	}
	return articles
}
