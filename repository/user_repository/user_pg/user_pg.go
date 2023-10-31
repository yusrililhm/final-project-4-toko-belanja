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

	updateUser = `
		UPDATE "users"
		SET 
			full_name = $2,
			email = $3,
			updated_at = now()
		WHERE id = $1
		RETURNING
			id, full_name, email, updated_at;
	`

	deleteUser = `
		DELETE FROM
			"users"
		WHERE
			id = $1;
	`

	adminQuery = `
		INSERT INTO users
			(
				full_name,
				email,
				password,
				role,
				balance
			)
		VALUES ('admin', 'admin@hacktivate.com', $1, 'admin', '0')
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

//Update User

func (u *userPg) UpdateUser(userPayLoad *entity.User) (*dto.UserUpdateResponse, errs.Error) {
	tx, err := u.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	row := tx.QueryRow(updateUser, userPayLoad.Id, userPayLoad.FullName, userPayLoad.Email)

	var userUpdate dto.UserUpdateResponse
	err = row.Scan(
		&userUpdate.Id,
		&userUpdate.FullName,
		&userUpdate.Email,
		&userUpdate.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &userUpdate, nil
}

//Delete User
func (u *userPg) DeleteUser(userId int) errs.Error {
	tx, err := u.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(deleteUser, userId)

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

// ============ Admin ==============
func (u *userPg) Admin(userPayLoad *entity.User) errs.Error {
	tx, err := u.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}
	_, err = tx.Exec(adminQuery, userPayLoad.Password)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong " + err.Error())
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}