package dto

import (
	"time"
	"toko-belanja-app/entity"
)

type CategoriesRequest struct {
	Type string `json:"type" valid:"required~Type can't be empty" example:"jersey"`
}

type CreateNewCategoriesResponse struct {
	Id                int       `json:"id" example:"1"`
	Type              string    `json:"type" example:"jersey"`
	SoldProductAmount int       `json:"sold_product_amount" example:"0"`
	CreatedAt         time.Time `json:"created_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type UpdateCategoryResponse struct {
	Id                int       `json:"id" example:"1"`
	Type              string    `json:"type" example:"jersey"`
	SoldProductAmount int       `json:"sold_product_amount" example:"0"`
	UpdatedAt         time.Time `json:"updated_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type GetCategories struct {
	Id                int              `json:"id"`
	Type              string           `json:"type"`
	SoldProductAmount int              `json:"sold_product_amount"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
	Products          []entity.Product `json:"products"`
}

type CategoryResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
