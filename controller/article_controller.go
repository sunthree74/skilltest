package controller

import (
	"net/http"
	"strconv"
	"sunthree74/skilltest/helper"
	"sunthree74/skilltest/model/web"
	"sunthree74/skilltest/service"

	"github.com/julienschmidt/httprouter"
)

type ArticleController struct {
	ArticleServiceInterface service.ArticleServiceInterface
}

func NewArticleControllerInterface(articleServiceInterface service.ArticleServiceInterface) ArticleControllerInterface {
	return &ArticleController{
		ArticleServiceInterface: articleServiceInterface,
	}
}

func (controller *ArticleController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	articleCreateRequest := web.ArticleCreateRequest{}
	helper.ReadFromRequestBody(request, &articleCreateRequest)

	articleResponse := controller.ArticleServiceInterface.Create(request.Context(), articleCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ArticleController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	articleUpdateRequest := web.ArticleUpdateRequest{}
	helper.ReadFromRequestBody(request, &articleUpdateRequest)

	articleId := params.ByName("articleId")
	id, err := strconv.Atoi(articleId)
	helper.PanicIfError(err)

	articleUpdateRequest.Id = id

	articleResponse := controller.ArticleServiceInterface.Update(request.Context(), articleUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ArticleController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	articleId := params.ByName("articleId")
	id, err := strconv.Atoi(articleId)
	helper.PanicIfError(err)

	controller.ArticleServiceInterface.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ArticleController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	articleId := params.ByName("articleId")
	id, err := strconv.Atoi(articleId)
	helper.PanicIfError(err)

	articleResponse := controller.ArticleServiceInterface.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ArticleController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	QueryTitle := request.URL.Query().Get("title")
	QueryCategoryName := request.URL.Query().Get("category_name")
	articleResponses := controller.ArticleServiceInterface.FindAll(request.Context(), QueryTitle, QueryCategoryName)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   articleResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
