package auth_service

import "github.com/gin-gonic/gin"

type AuthService interface {
	Authorization() gin.HandlerFunc
}

type authServiceImpl struct {
}

func NewAuthService() AuthService {
	return &authServiceImpl{}
}

// Authorization implements AuthService.
func (a *authServiceImpl) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// auth logic here
	}
}
