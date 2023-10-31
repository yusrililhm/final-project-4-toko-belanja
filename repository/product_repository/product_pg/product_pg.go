package product_pg

import (
	"database/sql"
	"toko-belanja-app/repository/product_repository"
)

type productPg struct {
	db *sql.DB
}

func NewProductPg(db *sql.DB) product_repository.ProductRepository {
	return &productPg{db: db}
}
