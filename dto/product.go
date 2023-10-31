package dto

import "time"

type ProductRequest struct {
	Title      string `json:"title" valid:"required~Title can't be empty" example:"Jersey King MU 2023/2024"`
	Price      uint   `json:"price" valid:"required~Price can't be empty, range(0|50000000)~" example:"120000"`
	Stock      uint   `json:"stock" valid:"required~Stock can't be empty, range(5|1000000)~" example:"10"`
	CategoryId int    `json:"category_id" example:"1"`
}

type GetProduct struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      uint      `json:"stock"`
	CategoryId int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type NewProductResponse struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      uint      `json:"stock"`
	CategoryId int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type UpdateProductResponse struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      uint      `json:"stock"`
	CategoryId int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
