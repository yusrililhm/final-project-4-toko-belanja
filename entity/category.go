package entity

import "time"

type Category struct {
	Id                int
	Type              string
	SoldProductAmount int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
