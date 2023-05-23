package dto

type CategoryCreateRequestDto struct {
	Name string `validate:"required,min=3,max=200"`
}
