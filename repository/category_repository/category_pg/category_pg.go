package category_pg

import (
	"database/sql"
	"toko-belanja-app/repository/category_repository"
)

type categoryPg struct {
	db *sql.DB
}

func NewCategoryPg(db *sql.DB) category_repository.CategoryRepository {
	return &categoryPg{db: db}
}
