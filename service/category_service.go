package service

import (
	"belajar_golang_api/model/dto"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request dto.CategoryCreateRequestDto) dto.CategoryResponse
	Update(ctx context.Context, request dto.CategoryUpdateRequestDto) dto.CategoryResponse
	Delete(ctx context.Context, categoryId int64)
	FindById(ctx context.Context, categoryId int64) dto.CategoryResponse
	FindAll(ctx context.Context) []dto.CategoryResponse
}
