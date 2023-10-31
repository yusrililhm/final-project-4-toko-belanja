package handler

import (
	"toko-belanja-app/service/user_service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UserRegister(ctx *gin.Context)
	UserLogin(ctx *gin.Context)
	UserTopUp(ctx *gin.Context)
}

type userHandlerImpl struct {
	us user_service.UserService
}

func NewUserHandler(userService user_service.UserService) UserHandler {
	return &userHandlerImpl{us: userService}
}

// UserLogin implements UserHandler.
func (uh *userHandlerImpl) UserLogin(ctx *gin.Context) {

}

// UserRegister implements UserHandler.
func (uh *userHandlerImpl) UserRegister(ctx *gin.Context) {

}

// UserTopUp implements UserHandler.
func (uh *userHandlerImpl) UserTopUp(ctx *gin.Context) {

}
