package server

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.Default()
	teamController := getTeamController()

	router.GET("/team/:id", teamController.Get)
	router.POST("/load-teams", teamController.Load)
	return router

}
