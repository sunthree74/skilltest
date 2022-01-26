package helper

import (
	"sunthree74/skilltest/model/entity"
	"sunthree74/skilltest/model/web"
)

func ToArticleResponse(article entity.Article) web.ArticleResponse {
	return web.ArticleResponse{
		Id:         article.Id,
		Title:      article.Title,
		Slug:       article.Slug,
		CategoryId: article.CategoryId,
		Content:    article.Content,
		Category:   article.Category,
		CreatedAt:  article.CreatedAt,
		UpdatedAt:  article.UpdatedAt,
	}
}

func ToArticleResponses(articles []entity.Article) []web.ArticleResponse {
	var articleResponses []web.ArticleResponse
	for _, article := range articles {
		articleResponses = append(articleResponses, ToArticleResponse(article))
	}
	return articleResponses
}

func ToCategoryResponse(category entity.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:           category.Id,
		CategoryName: category.CategoryName,
		CategorySlug: category.CategorySlug,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}
}

func ToCategoryResponses(categories []entity.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
