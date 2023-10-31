package handler

import (
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
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
// UserLogin godoc
// @Summary User login
// @Description User login
// @Tags Users
// @Accept json
// @Produce json
// @Param dto.UsersLoginRequest body dto.UsersLoginRequest true "body request for user login"
// @Success 200 {object} dto.UserResponse
// @Router /users/login [post]
func (uh *userHandlerImpl) UserLogin(ctx *gin.Context) {

	loginRequest := &dto.UsersLoginRequest{}

	if err := ctx.ShouldBindJSON(loginRequest); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}
}

// UserRegister implements UserHandler.
// UserRegister godoc
// @Summary User register
// @Description User register
// @Tags Users
// @Accept json
// @Produce json
// @Param dto.CreateNewUsersRequest body dto.CreateNewUsersRequest true "body request for user register"
// @Success 201 {object} dto.UserResponse
// @Router /users/register [post]
func (uh *userHandlerImpl) UserRegister(ctx *gin.Context) {

	registeRequest := &dto.CreateNewUsersRequest{}

	if err := ctx.ShouldBindJSON(registeRequest); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}
}

// UserTopUp implements UserHandler.
// UserTopUp godoc
// @Summary User topup
// @Description User topup
// @Tags Users
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param dto.UsersTopUpRequest body dto.UsersTopUpRequest true "body request for user topup"
// @Success 200 {object} dto.UserResponse
// @Router /users/topup [patch]
func (uh *userHandlerImpl) UserTopUp(ctx *gin.Context) {

	topupRequest := &dto.UsersTopUpRequest{}

	if err := ctx.ShouldBindJSON(topupRequest); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	user := ctx.MustGet("userData").(entity.User)
	_ = user.Id
}
