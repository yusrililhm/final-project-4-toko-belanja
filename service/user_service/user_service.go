package user_service

import "toko-belanja-app/repository/user_repository"

type UserService interface {
}

type userServiceImpl struct {
	ur user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userServiceImpl{ur: userRepo}
}
