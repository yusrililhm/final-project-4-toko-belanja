package user_repository

import (
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(userPayLoad *entity.User) (*dto.CreateNewUsersResponse, errs.Error)
}
