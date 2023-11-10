package category_pg

import (
	"database/sql"
	"time"
	"toko-belanja-app/entity"
	"toko-belanja-app/repository/category_repository"
)

type categoryWithProduct struct {
	CategoryId                int
	CategoryType              string
	CategorySoldProductAmount int
	CategoryCreatedAt         time.Time
	CategoryUpdatedAt         time.Time
	ProductId                 sql.NullInt64
	ProductTitle              sql.NullString
	ProductPrice              sql.NullInt64
	ProductStock              sql.NullInt64
	ProductCreatedAt          sql.NullTime
	ProductUpdatedAt          sql.NullTime
}

func (c *categoryWithProduct) categoryWithProductToEntity() category_repository.CategoryProduct {
	return category_repository.CategoryProduct{
		Category: entity.Category{
			Id:                c.CategoryId,
			Type:              c.CategoryType,
			SoldProductAmount: c.CategorySoldProductAmount,
			CreatedAt:         c.CategoryCreatedAt,
			UpdatedAt:         c.CategoryUpdatedAt,
		},
		Product: entity.Product{
			Id:        int(c.ProductId.Int64),
			Title:     c.ProductTitle.String,
			Price:     uint(c.ProductPrice.Int64),
			Stock:     uint(c.ProductStock.Int64),
			CreatedAt: c.ProductCreatedAt.Time,
			UpdatedAt: c.ProductUpdatedAt.Time,
		},
	}
}
