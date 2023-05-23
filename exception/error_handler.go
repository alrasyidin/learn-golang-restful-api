package exception

import (
	"belajar_golang_api/helper"
	"belajar_golang_api/model/dto"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorPanicHandler(w http.ResponseWriter, r *http.Request, error any) {
	if notFoundErrorHanler(w, r, error) {
		return
	}

	if badRequestErrorHanler(w, r, error) {
		return
	}

	internalServerErrorHandler(w, r, error)
}

func badRequestErrorHanler(w http.ResponseWriter, r *http.Request, error any) bool {
	exception, ok := error.(validator.ValidationErrors)
	fmt.Printf("validation %+v", exception)
	fmt.Printf("ok %+v", ok)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundErrorHanler(w http.ResponseWriter, r *http.Request, error any) bool {
	exception, ok := error.(NotFoundError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := dto.WebResponse{
			Code:   http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := dto.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}
