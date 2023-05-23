package dto

type CategoryUpdateRequestDto struct {
	Id   int64  `validate:"required"`
	Name string `validate:"required,min=3,max=200"`
}
