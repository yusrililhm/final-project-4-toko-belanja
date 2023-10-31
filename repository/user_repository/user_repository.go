package user_repository

import (
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(userPayLoad *entity.User) (*dto.CreateNewUsersResponse, errs.Error)
	TopUpBalance(userPayLoad *entity.User) (*dto.TopUpResponse, errs.Error)
	GetUserByEmail(userEmail string) (*entity.User, errs.Error)
	GetUserById(userId int) (*entity.User, errs.Error)
}
