// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"fmt"
	"net/http"

	"tpc/docs"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RouterV1 -.
type RouterV1 struct {
	g *gin.RouterGroup
}

// NewRouter -.
// @title       TOC PD CAPTAIN
// @description API docs for Auto Trade
// @version     1.0.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func NewRouter(handler *gin.Engine) *RouterV1 {
	prefix := "/tpc/v1"

	docs.SwaggerInfo.BasePath = prefix
	docs.SwaggerInfo.Host = "127.0.0.1:16888"

	// handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Prometheus metrics
	handler.GET(fmt.Sprintf("%s/-/health", prefix), healthCheck)

	return &RouterV1{
		g: handler.Group(prefix),
	}
}

// @Summary     healthCheck
// @Description healthCheck
// @ID          healthCheck
// @Tags  	    healthCheck
// @Accept      json
// @Produce     json
// @Success     200 {string} string
// @Router      /-/health [get]
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
