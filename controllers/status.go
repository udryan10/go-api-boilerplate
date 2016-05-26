package controllers

import (
	"github.com/gin-gonic/gin"
)

// Status check
func Status(c *gin.Context) {
	// return 200 with empty JSON body
	c.JSON(200, gin.H{})
}
