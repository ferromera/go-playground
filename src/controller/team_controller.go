package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
}

func (tc *TeamController) Get(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "This is the team %s", name)
}
