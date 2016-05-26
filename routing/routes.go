package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/udryan10/go-api-boilerplate/controllers"
)

// Handlers for route management
func Handlers(router *gin.Engine) {
	// GET /status
	router.GET("/status", controllers.Status)
}
