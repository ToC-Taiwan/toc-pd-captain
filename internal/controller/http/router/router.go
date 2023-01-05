// Package router implements routing paths. Each services in own file.
package router

import (
	"fmt"
	"net/http"

	"tpc/docs"
	"tpc/internal/usecase"
	"tpc/pkg/log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = log.New()

const (
	urlPrefix = "/tpc"
)

// RouterV1 -.
type RouterV1 struct {
	public    *gin.RouterGroup
	privateV1 *gin.RouterGroup
}

// NewRouter -.
// @title       ToC PD Capatain
// @description API doc for PD Capatain
// @version     1.0.0-alpha
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func NewRouter(handler *gin.Engine, userRepo usecase.User) *RouterV1 {
	docs.SwaggerInfo.BasePath = urlPrefix

	// handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(corsMiddleware())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/doc/*any", swaggerHandler)

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// new auth
	auth := newAuthMiddleware(userRepo)

	// public
	public := handler.Group(urlPrefix)
	public.GET("/-/health", healthCheck)
	public.POST("/login", loginHandler(auth))

	privateV1 := handler.Group(fmt.Sprintf("%s/v1", urlPrefix))
	privateV1.Use(auth.MiddlewareFunc())
	privateV1.Use(tokenInterceptor(auth))

	return &RouterV1{
		public:    public,
		privateV1: privateV1,
	}
}

// @Summary     healthCheck
// @Description healthCheck
// @ID          healthCheck
// @Tags  	    Health
// @Accept      json
// @Produce     json
// @Success     200 {string} string
// @Router      /-/health [get]
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

// @Summary     Login
// @Description Every api request will extend token expired time, websocket will not extend.
// @tags Login V1
// @accept json
// @produce json
// @param body body entity.User{} true "Body"
// @success 200 {object} loginResponseBody{}
// @router /login [post]
func loginHandler(mid *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return mid.LoginHandler
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		docs.SwaggerInfo.Host = c.Request.Host
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Set("content-type", "application/json")
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, nil)
		}
		c.Next()
	}
}

func tokenInterceptor(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, err := authMiddleware.RefreshToken(c)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Next()
	}
}
