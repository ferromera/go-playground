package server

import(
	"github.com/gin-gonic/gin"
)
func New() *gin.Engine {
	router := gin.Default()
	teamController := getTeamController()
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/team/:id", teamController.Get)
	return router

}