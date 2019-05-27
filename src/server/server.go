package server

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.Default()
	teamController := getTeamController()
	playerController := getPlayerController()

	router.GET("/teams/:id", teamController.Get)
	router.GET("/teams", teamController.GetAll)
	router.GET("/teams/:id/players", teamController.GetPlayers)
	router.POST("/load-teams", teamController.Load)

	router.GET("/players/:id", playerController.Get)
	router.GET("/players", playerController.GetFilter)
	router.POST("/load-players", playerController.Load)

	return router

}
