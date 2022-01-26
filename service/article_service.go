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

type ArticleService struct {
	ArticleRepositoryInterface repository.ArticleRepositoryInterface
	DB                         *sql.DB
	Validate                   *validator.Validate
}

func NewArticleServiceInterface(articleRepositoryInterface repository.ArticleRepositoryInterface, DB *sql.DB, validate *validator.Validate) ArticleServiceInterface {
	return &ArticleService{
		ArticleRepositoryInterface: articleRepositoryInterface,
		DB:                         DB,
		Validate:                   validate,
	}
}

func (service *ArticleService) Create(ctx context.Context, request web.ArticleCreateRequest) web.ArticleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article := entity.Article{
		Title:      request.Title,
		Slug:       request.Slug,
		CategoryId: request.CategoryId,
		Content:    request.Content,
	}

	article = service.ArticleRepositoryInterface.Save(ctx, tx, article)

	return helper.ToArticleResponse(article)
}

func (service *ArticleService) Update(ctx context.Context, request web.ArticleUpdateRequest) web.ArticleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, err := service.ArticleRepositoryInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	article.Title = request.Title
	article.Slug = request.Slug
	article.CategoryId = request.CategoryId
	article.Content = request.Content

	article = service.ArticleRepositoryInterface.Update(ctx, tx, article)

	return helper.ToArticleResponse(article)
}

func (service *ArticleService) Delete(ctx context.Context, articleId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, err := service.ArticleRepositoryInterface.FindById(ctx, tx, articleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ArticleRepositoryInterface.Delete(ctx, tx, article)
}

func (service *ArticleService) FindById(ctx context.Context, articleId int) web.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	article, err := service.ArticleRepositoryInterface.FindById(ctx, tx, articleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToArticleResponse(article)
}

func (service *ArticleService) FindAll(ctx context.Context, QueryTitle string, QueryCategoryName string) []web.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	articles := service.ArticleRepositoryInterface.FindAll(ctx, tx, QueryTitle, QueryCategoryName)

	return helper.ToArticleResponses(articles)
}
