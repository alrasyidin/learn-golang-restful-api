package dto

type CategoryResponse struct {
	Id   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
