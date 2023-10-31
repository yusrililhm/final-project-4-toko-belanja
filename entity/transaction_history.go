package entity

import "time"

type TransactionHistory struct {
	Id         int
	UserId     int
	ProductId  int
	Quantity   uint
	TotalPrice uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
