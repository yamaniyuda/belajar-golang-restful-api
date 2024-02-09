package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/service"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller *CategoryControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) Update(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(err)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) FindById(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
