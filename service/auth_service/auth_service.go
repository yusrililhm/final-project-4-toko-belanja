package auth_service

import (
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
}

type authServiceImpl struct {
	ur user_repository.UserRepository
}

func NewAuthService(ur user_repository.UserRepository) AuthService {
	return &authServiceImpl{
		ur: ur,
	}
}

// Authentication Implements AuthService.
func (a *authServiceImpl) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.ur.GetUserById(user.Id)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_, err = a.ur.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		ctx.Set("userData", user)
		ctx.Next()
	}
}

// Authorization implements AuthService.
func (a *authServiceImpl) AdminAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(entity.User)
		if !ok {
			newError := errs.NewBadRequestError("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}
		if userData.Role != "admin" {
			newError := errs.NewUnathorizedError("You're not authorized to access this endpoint")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		ctx.Next()
	}
}
