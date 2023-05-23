package controller

import (
	"belajar_golang_api/helper"
	"belajar_golang_api/model/dto"
	"belajar_golang_api/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	Service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		Service: service,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := dto.CategoryCreateRequestDto{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.Service.Create(r.Context(), categoryCreateRequest)
	webResponse := dto.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateRequest := dto.CategoryUpdateRequestDto{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.ParseInt(categoryId, 10, 64)
	helper.HandleIfPanicError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.Service.Update(r.Context(), categoryUpdateRequest)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.ParseInt(categoryId, 10, 64)
	helper.HandleIfPanicError(err)

	controller.Service.Delete(r.Context(), id)

	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.ParseInt(categoryId, 10, 64)
	helper.HandleIfPanicError(err)

	categoryResponse := controller.Service.FindById(r.Context(), id)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("\nini kesini")
	categoryResponses := controller.Service.FindAll(r.Context())
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
