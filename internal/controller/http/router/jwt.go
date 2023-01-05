package router

import (
	"net/http"
	"time"

	"tpc/internal/entity"
	"tpc/internal/usecase"
	"tpc/pkg/utils"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	realm = "ToC"

	identityKey = "toc_pd_captain"
	cookieName  = "toc_pdc_token"

	tokenHeadName = "Bearer"
	tokenLookup   = "header: Authorization, query: token, cookie: token"
)

func newAuthMiddleware(userRepo usecase.User) *jwt.GinJWTMiddleware {
	auth := newUserAuthenticator(userRepo)
	j, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:          realm,
		Key:            []byte(utils.RandomString(20)),
		SendCookie:     true,
		SecureCookie:   false,
		CookieHTTPOnly: true,
		CookieName:     cookieName,
		CookieSameSite: http.SameSiteDefaultMode,
		Timeout:        time.Hour,
		MaxRefresh:     time.Hour,
		IdentityKey:    identityKey,
		TokenLookup:    tokenLookup,
		TokenHeadName:  tokenHeadName,
		TimeFunc:       time.Now,

		PayloadFunc:     payloadFunc,
		LoginResponse:   loginResponse,
		IdentityHandler: identityHandler,
		Unauthorized:    unauthorized,
		Authenticator:   auth.authenticator,
		Authorizator:    auth.authorizator,
	})
	if err != nil {
		logger.Fatal(err)
	}
	return j
}

type loginResponseBody struct {
	Token  string `json:"token" yaml:"token"`
	Expire string `json:"expire" yaml:"expire"`
}

func loginResponse(c *gin.Context, code int, message string, currentTime time.Time) {
	c.JSON(http.StatusOK, loginResponseBody{
		Token:  message,
		Expire: currentTime.Format(time.RFC3339),
	})
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*entity.User); ok {
		return jwt.MapClaims{
			identityKey: v.UserName,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &entity.User{
		UserName: claims[identityKey].(string),
	}
}

func unauthorized(c *gin.Context, code int, message string) {
	ErrorResponse(c, code, message)
}
