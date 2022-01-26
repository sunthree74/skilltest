package main

import (
	"net/http"
	"sunthree74/skilltest/app"
	"sunthree74/skilltest/controller"
	"sunthree74/skilltest/helper"
	"sunthree74/skilltest/middleware"
	"sunthree74/skilltest/repository"
	"sunthree74/skilltest/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepositoryInterface()
	categoryService := service.NewCategoryServiceInterface(categoryRepository, db, validate)
	categoryController := controller.NewCategoryControllerInterface(categoryService)

	articleRepository := repository.NewArticleRepositoryInterface()
	articleService := service.NewArticleServiceInterface(articleRepository, db, validate)
	articleController := controller.NewArticleControllerInterface(articleService)
	router := app.NewRouter(articleController, categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
