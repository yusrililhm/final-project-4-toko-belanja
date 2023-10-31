package user_pg

import (
	"database/sql"
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/repository/user_repository"
)

const (
	createNewUser = `
		INSERT INTO "users"
		(
			full_name,
			email,
			password,
			role,
			balance
		)
		VALUES ($1, $2, $3, 'customer', '0')
		RETURNING
			id, full_name, email, password, balance, created_at;
	`
	UsersTopUp = `
		UPDATE "users"
		SET balance = balance + $2
		WHERE id = $1
	`
	)

	

type userPg struct {
	db *sql.DB
}

func NewUserPg(db *sql.DB) user_repository.UserRepository {
	return &userPg{db: db}
}

//Create New User
func (u *userPg) CreateNewUser(userPayLoad *entity.User) (*dto.CreateNewUsersResponse, errs.Error) {
	tx, err := u.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	var user dto.CreateNewUsersResponse
	row := tx.QueryRow(createNewUser, userPayLoad.FullName, userPayLoad.Email, userPayLoad.Password)

	err = row.Scan(&user.Id, &user.FullName, &user.Email, &user.CreatedAt)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &user, nil
}


// Top Up
func (u *userPg) TopUpBalance(userPayLoad *entity.User) (*dto.UsersTopUpRequest, errs.Error) {
    tx, err := u.db.Begin()

    if err != nil {
        tx.Rollback()
        return nil, errs.NewInternalServerError("something went wrong")
    }

    var user dto.UsersTopUpRequest
    row := tx.QueryRow(UsersTopUp, userPayLoad.Id, userPayLoad.Balance)

    err = row.Scan(&user.Balance)

    if err != nil {
        tx.Rollback()
        return nil, errs.NewInternalServerError("something went wrong")
    }

    err = tx.Commit()

    if err != nil {
        tx.Rollback()
        return nil, errs.NewInternalServerError("something went wrong")
    }
    return &user, nil
}

