package service

import (
	"context"
	"sunthree74/skilltest/model/web"
)

type ArticleServiceInterface interface {
	Create(ctx context.Context, request web.ArticleCreateRequest) web.ArticleResponse
	Update(ctx context.Context, request web.ArticleUpdateRequest) web.ArticleResponse
	Delete(ctx context.Context, articleId int)
	FindById(ctx context.Context, articleId int) web.ArticleResponse
	FindAll(ctx context.Context, QueryTitle string, QueryCategoryName string) []web.ArticleResponse
}
