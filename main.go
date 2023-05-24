package main

import (
	"belajar_golang_api/helper"
	"belajar_golang_api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	server := InitializedServer()

	err := server.ListenAndServe()
	helper.HandleIfPanicError(err)
	// fmt.Printf("Server running on http://localhost:%v", PORT)
	// done := make(chan bool)
	// <-done
}

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}
