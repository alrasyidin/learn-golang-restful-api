package middleware

import (
	"belajar_golang_api/helper"
	"belajar_golang_api/model/dto"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

// Constructor for AuthMiddleware
func NewAuthMiddleware(Handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: Handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-Key") == "RAHASIA" {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := dto.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
