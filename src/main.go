package main

import (
	"github.com/ferromera/go-playground/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	teamController := &controller.TeamController{}
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/team/:name", teamController.Get)

	router.Run(":8080")
}
