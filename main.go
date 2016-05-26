package main

import (
	"github.com/gin-gonic/gin"
	"github.com/udryan10/go-api-boilerplate/routing"
)

func main() {
	// initialize gin
	router := gin.Default()
	// initialize router handlers
	routing.Handlers(router)
	router.Run(":8080")
}
