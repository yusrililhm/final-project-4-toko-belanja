package user_pg

import (
	"database/sql"
	"errors"
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
		VALUES ($1, $2, $3, 'customer', 0)
		RETURNING
			id, full_name, email, password, balance, created_at;
	`
	UsersTopUp = `
		UPDATE "users"
		SET balance = balance + $2
		WHERE id = $1
	`

	getUserById = `
		SELECT
			id, 
			full_name, 
			email, 
			password,
			balance, 
			role, 
			created_at, 
			updated_at
		FROM "users"
		WHERE 
			id = $1;
	`
	getUserByEmail = `
			SELECT
				id, 
				full_name, 
				email, 
				password,
				balance, 
				role, 
				created_at, 
				updated_at
			FROM "users"
			WHERE email = $1;
		`
)

type userPg struct {
	db *sql.DB
}

func NewUserPg(db *sql.DB) user_repository.UserRepository {
	return &userPg{db: db}
}

// Create New User
func (u *userPg) CreateNewUser(userPayLoad *entity.User) (*dto.CreateNewUsersResponse, errs.Error) {
	tx, err := u.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	var user dto.CreateNewUsersResponse
	row := tx.QueryRow(createNewUser, userPayLoad.FullName, userPayLoad.Email, userPayLoad.Password, userPayLoad.Balance)

	err = row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Balance, &user.CreatedAt)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}
	
	return &user, nil
}

// Top Up
func (u *userPg) TopUpBalance(userPayLoad *entity.User) errs.Error {
	tx, err := u.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(UsersTopUp, &userPayLoad.Id, &userPayLoad.Balance)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (u *userPg) GetUserById(userId int) (*entity.User, errs.Error) {
	var user entity.User

	row := u.db.QueryRow(getUserById, userId)

	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Balance, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	return &user, nil
}

func (u *userPg) GetUserByEmail(email string) (*entity.User, errs.Error) {
	var user entity.User

	row := u.db.QueryRow(getUserByEmail, email)

	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Balance, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	return &user, nil
}
