package dto

import (
	"time"
	"toko-belanja-app/entity"
)

type TransactionRequest struct {
	ProductId int  `json:"product_id" example:"1"`
	Quantity  uint `json:"quantity" valid:"required~Quantity can't be empty" example:"3"`
}

type TransactionBill struct {
	TotalPrice   uint   `json:"total_price"`
	Quantity     uint   `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type TransactionResponse struct {
	TransactionBill TransactionBill `json:"transaction_bill"`
}

type User struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Balance   uint      `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MyTransaction struct {
	Id         int            `json:"id"`
	UserId     int            `json:"user_id"`
	ProductId  int            `json:"product_id"`
	Quantity   uint           `json:"quantity"`
	TotalPrice uint           `json:"total_price"`
	Products   entity.Product `json:"products"`
}

type UsersTransaction struct {
	Id         int            `json:"id"`
	UserId     int            `json:"user_id"`
	ProductId  int            `json:"product_id"`
	Quantity   uint           `json:"quantity"`
	TotalPrice uint           `json:"total_price"`
	Products   entity.Product `json:"products"`
	User       User           `json:"users"`
}

type TransactionHistoryResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
