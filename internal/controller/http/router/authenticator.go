package router

import (
	"tpc/internal/entity"
	"tpc/internal/usecase"
	"tpc/pkg/utils"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type userAuthenticator struct {
	uc usecase.User
}

func newUserAuthenticator(uc usecase.User) *userAuthenticator {
	return &userAuthenticator{uc}
}

func (u *userAuthenticator) authenticator(c *gin.Context) (interface{}, error) {
	var loginVals entity.User
	if err := c.ShouldBind(&loginVals); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	user := loginVals.UserName
	password := loginVals.Password

	dbUser, err := u.uc.GetUserByUserName(user)
	if err != nil {
		return nil, err
	}

	if user == dbUser.UserName && utils.SHA256Generator(password) == dbUser.Password {
		return &entity.User{UserName: user}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func (u *userAuthenticator) authorizator(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*entity.User); ok {
		return true
	}
	return false
}
