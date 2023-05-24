package service

import (
	"belajar_golang_api/exception"
	"belajar_golang_api/helper"
	"belajar_golang_api/model/domain"
	"belajar_golang_api/model/dto"
	"belajar_golang_api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	DB                 *sql.DB
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		DB:                 db,
		CategoryRepository: categoryRepository,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request dto.CategoryCreateRequestDto) dto.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.HandleIfPanicError(err)

	tx, err := service.DB.Begin()
	helper.HandleIfPanicError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request dto.CategoryUpdateRequestDto) dto.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.HandleIfPanicError(err)

	tx, err := service.DB.Begin()
	helper.HandleIfPanicError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int64) {
	tx, err := service.DB.Begin()
	helper.HandleIfPanicError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int64) dto.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.HandleIfPanicError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, int64(categoryId))

	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []dto.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.HandleIfPanicError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}
