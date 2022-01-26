package service

import (
	"context"
	"database/sql"
	"sunthree74/skilltest/exception"
	"sunthree74/skilltest/helper"
	"sunthree74/skilltest/model/entity"
	"sunthree74/skilltest/model/web"
	"sunthree74/skilltest/repository"

	"github.com/go-playground/validator"
)

type CategoryService struct {
	CategoryRepositoryInterface repository.CategoryRepositoryInterface
	DB                          *sql.DB
	Validate                    *validator.Validate
}

func NewCategoryServiceInterface(categoryRepositoryInterface repository.CategoryRepositoryInterface, DB *sql.DB, validate *validator.Validate) CategoryServiceInterface {
	return &CategoryService{
		CategoryRepositoryInterface: categoryRepositoryInterface,
		DB:                          DB,
		Validate:                    validate,
	}
}

func (service *CategoryService) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := entity.Category{
		CategoryName: request.CategoryName,
		CategorySlug: request.CategorySlug,
	}

	category = service.CategoryRepositoryInterface.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryService) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepositoryInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.CategoryName = request.CategoryName
	category.CategorySlug = request.CategorySlug

	category = service.CategoryRepositoryInterface.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryService) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepositoryInterface.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepositoryInterface.Delete(ctx, tx, category)
}

func (service *CategoryService) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepositoryInterface.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryService) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepositoryInterface.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
