package user_service

import (
	"fmt"
	"net/http"
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/pkg/helpers"
	"toko-belanja-app/repository/user_repository"
)

type UserService interface {
	RegisterUser(userPayLoad *dto.CreateNewUsersRequest) (*dto.UserResponse, errs.Error)
	LoginUser(userPayLoad *dto.UsersLoginRequest) (*dto.UserResponse, errs.Error)
	TopUpBalance(userId int, userPayLoad *dto.UsersTopUpRequest) (*dto.UserResponse, errs.Error)
}

type userServiceImpl struct {
	ur user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userServiceImpl{ur: userRepo}
}

func (us *userServiceImpl) RegisterUser(userPayLoad *dto.CreateNewUsersRequest) (*dto.UserResponse, errs.Error) {
	err := helpers.ValidateStruct(userPayLoad)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		FullName: userPayLoad.FullName,
		Email:    userPayLoad.Email,
		Password: userPayLoad.Password,
	}

	user.HashPassword()

	response, err := us.ur.CreateNewUser(user)

	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		Code:    http.StatusCreated,
		Message: "Your account has been successfully created",
		Data:    response,
	}, nil
}

func (us *userServiceImpl) LoginUser(userPayLoad *dto.UsersLoginRequest) (*dto.UserResponse, errs.Error) {
	err := helpers.ValidateStruct(userPayLoad)

	if err != nil {
		return nil, err
	}

	user, err := us.ur.GetUserByEmail(userPayLoad.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(userPayLoad.Password)

	if isValidPassword == false {
		return nil, errs.NewBadRequestError("invalid email/password")
	}

	token := user.GenerateToken()

	return &dto.UserResponse{
		Code:    http.StatusOK,
		Message: "You have successfully logged into your account",
		Data: dto.UsersLoginResponse{
			Token: token,
		},
	}, nil
}

func (us *userServiceImpl) TopUpBalance(userId int, userPayLoad *dto.UsersTopUpRequest) (*dto.UserResponse, errs.Error) {
	err := helpers.ValidateStruct(userPayLoad)

	if err != nil {
		return nil, err
	}

	user, err := us.ur.GetUserById(userId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("invalid user")
		}
		return nil, err
	}

	if user.Id != userId {
		return nil, errs.NewNotFoundError("invalid user")
	}

	usr := &entity.User{
		Id:      userId,
		Balance: userPayLoad.Balance,
	}

	response, err := us.ur.TopUpBalance(usr)

	if err != nil {
		return nil, err
	}

	balance := fmt.Sprintf("Your balance has been successfully updated to Rp%d", response.Balance)

	return &dto.UserResponse{
		Code:    http.StatusOK,
		Message: balance,
		Data:    nil,
	}, nil
}
