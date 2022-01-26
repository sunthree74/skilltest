package app

import (
	"sunthree74/skilltest/controller"
	"sunthree74/skilltest/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(articleController controller.ArticleControllerInterface, categoryController controller.CategoryControllerInterface) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/articles/list", articleController.FindAll)
	router.GET("/api/article/detail/:articleId", articleController.FindById)
	router.POST("/api/article/create", articleController.Create)
	router.PUT("/api/article/update/:articleId", articleController.Update)
	router.DELETE("/api/article/delete/:articleId", articleController.Delete)

	router.GET("/api/categories/list", categoryController.FindAll)
	router.GET("/api/categorie/detail/:categoryId", categoryController.FindById)
	router.POST("/api/categorie/create", categoryController.Create)
	router.PUT("/api/categorie/update/:categoryId", categoryController.Update)
	router.DELETE("/api/categorie/delete/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
