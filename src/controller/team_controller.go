package controller

import (
	"net/http"

	"github.com/ferromera/go-playground/src/domain"
	"github.com/gin-gonic/gin"
)

type TeamController struct {
	Service domain.TeamResolver
}

func (tc *TeamController) Get(c *gin.Context) {
	id := c.Param("id")
	team,_ := tc.Service.GetTeam(c,id)
	c.JSON( http.StatusOK, team)
}
