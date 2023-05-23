package main

import (
	"belajar_golang_api/app"
	"belajar_golang_api/controller"
	"belajar_golang_api/exception"
	"belajar_golang_api/helper"
	"belajar_golang_api/middleware"
	"belajar_golang_api/repository"
	"belajar_golang_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryContoller := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryContoller.FindAll)
	router.POST("/api/categories", categoryContoller.Create)
	router.GET("/api/categories/:categoryId", categoryContoller.FindById)
	router.PUT("/api/categories/:categoryId", categoryContoller.Update)
	router.DELETE("/api/categories/:categoryId", categoryContoller.Delete)

	// handling error
	router.PanicHandler = exception.ErrorPanicHandler

	// const PORT = 3000
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.HandleIfPanicError(err)
	// fmt.Printf("Server running on http://localhost:%v", PORT)
	// done := make(chan bool)
	// <-done
}
