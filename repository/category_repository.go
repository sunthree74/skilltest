package repository

import (
	"context"
	"database/sql"
	"errors"
	"sunthree74/skilltest/helper"
	"sunthree74/skilltest/model/entity"
	"time"
)

type CategoryRepository struct {
}

func NewCategoryRepositoryInterface() CategoryRepositoryInterface {
	return &CategoryRepository{}
}

func (repository *CategoryRepository) Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	SQL := "insert into category(category_name, category_slug, created_at, updated_at) values (?,?,?,?)"
	now := time.Now()
	now = now.Local()
	dateNow := now.Format("2006-01-02 15:04:05")
	result, err := tx.ExecContext(ctx, SQL, category.CategoryName, category.CategorySlug, dateNow, dateNow)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepository) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	SQL := "update category set category_name = ?, category_slug = ?, updated_at = ? where id = ? and deleted_at is null"
	now := time.Now()
	now = now.Local()
	dateNow := now.Format("2006-01-02 15:04:05")
	_, err := tx.ExecContext(ctx, SQL, category.CategoryName, category.CategorySlug, dateNow, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepository) Delete(ctx context.Context, tx *sql.Tx, category entity.Category) {
	SQL := "update category set deleted_at = ?  where id = ? and deleted_at is null"
	now := time.Now()
	now = now.Local()
	dateNow := now.Format("2006-01-02 15:04:05")
	_, err := tx.ExecContext(ctx, SQL, dateNow, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepository) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error) {
	SQL := "select id, category_name, category_slug, created_at, updated_at from category where id = ? and deleted_at is null"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := entity.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.CategoryName, &category.CategorySlug, &category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepository) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	SQL := "select id, category_name, category_slug, created_at, updated_at from category where deleted_at is null"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.Id, &category.CategoryName, &category.CategorySlug,
			&category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
