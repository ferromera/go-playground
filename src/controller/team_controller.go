package controller

import (
	"net/http"

	"github.com/ferromera/go-playground/src/domain"
	"github.com/ferromera/go-playground/src/errors"
	"github.com/gin-gonic/gin"
)

type TeamController struct {
	Service domain.TeamService
}

func (tc *TeamController) Get(c *gin.Context) {
	id := c.Param("id")
	team, err := tc.Service.GetTeam(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ApiError{
			Message:  err.Error(),
			ErrorStr: http.StatusText(http.StatusInternalServerError),
			Status:   http.StatusInternalServerError,
			Cause:    err.Error(),
		})
		return
	}
	if team == nil {
		c.JSON(http.StatusNotFound, errors.ApiError{
			Message:  "Team not found",
			ErrorStr: http.StatusText(http.StatusNotFound),
			Status:   http.StatusNotFound,
			Cause:    "Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, team)
}
func (tc *TeamController) GetAll(c *gin.Context) {

	teams, err := tc.Service.GetAllTeams(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ApiError{
			Message:  err.Error(),
			ErrorStr: http.StatusText(http.StatusInternalServerError),
			Status:   http.StatusInternalServerError,
			Cause:    err.Error(),
		})
		return
	}
	if teams == nil {
		c.JSON(http.StatusNotFound, errors.ApiError{
			Message:  "Teams not found",
			ErrorStr: http.StatusText(http.StatusNotFound),
			Status:   http.StatusNotFound,
			Cause:    "Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, teams)
}

func (tc *TeamController) GetPlayers(c *gin.Context) {
	id := c.Param("id")
	players, err := tc.Service.GetPlayers(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ApiError{
			Message:  err.Error(),
			ErrorStr: http.StatusText(http.StatusInternalServerError),
			Status:   http.StatusInternalServerError,
			Cause:    err.Error(),
		})
		return
	}
	if players == nil {
		c.JSON(http.StatusNotFound, errors.ApiError{
			Message:  "Players not found",
			ErrorStr: http.StatusText(http.StatusNotFound),
			Status:   http.StatusNotFound,
			Cause:    "Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, players)
}

func (tc *TeamController) Load(c *gin.Context) {

	count, err := tc.Service.Load(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ApiError{
			Message:  err.Error(),
			ErrorStr: http.StatusText(http.StatusInternalServerError),
			Status:   http.StatusInternalServerError,
			Cause:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, count)

}
