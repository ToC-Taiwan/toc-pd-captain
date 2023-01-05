package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type response struct {
	Response string `json:"response"`
}

// ErrorResponse -.
func ErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{
		Response: fmt.Sprintf("Error(%d): %s", code, msg),
	})
}
