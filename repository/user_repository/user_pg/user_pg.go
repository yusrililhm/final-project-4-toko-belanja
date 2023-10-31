package user_pg

import (
	"database/sql"
	"toko-belanja-app/repository/user_repository"
)

type userPg struct {
	db *sql.DB
}

func NewUserPg(db *sql.DB) user_repository.UserRepository {
	return &userPg{db: db}
}
