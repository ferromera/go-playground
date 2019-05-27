package controller

import (
	"net/http"
	"strconv"

	"github.com/ferromera/go-playground/src/domain"
	"github.com/ferromera/go-playground/src/errors"
	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	Service domain.PlayerService
}

func (tc *PlayerController) Get(c *gin.Context) {
	id := c.Param("id")
	player, err := tc.Service.GetPlayer(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ApiError{
			Message:  err.Error(),
			ErrorStr: http.StatusText(http.StatusInternalServerError),
			Status:   http.StatusInternalServerError,
			Cause:    err.Error(),
		})
		return
	}
	if player == nil {
		c.JSON(http.StatusNotFound, errors.ApiError{
			Message:  "Player not found",
			ErrorStr: http.StatusText(http.StatusNotFound),
			Status:   http.StatusNotFound,
			Cause:    "Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, player)
}
func (tc *PlayerController) validateFilter(c *gin.Context) error {
	minOverallQuery := c.Query("minOverall")
	maxOverallQuery := c.Query("maxOverall")
	var minOverall int
	var maxOverall int
	if minOverallQuery == "" {
		minOverall = 0
	}else{
		var err error
		minOverall, err = strconv.Atoi(minOverallQuery)
		if err != nil{
			return &errors.ApiError{
				Message:  "invalid minOverall",
				ErrorStr: http.StatusText(http.StatusBadRequest),
				Status:   http.StatusBadRequest,
				Cause:    err.Error(),
			}
		}
	}
	if maxOverallQuery == "" {
		maxOverall = 100
	}else{
		var err error
		maxOverall, err = strconv.Atoi(maxOverallQuery)
		if err != nil{
			return &errors.ApiError{
				Message:  "invalid maxOverall",
				ErrorStr: http.StatusText(http.StatusBadRequest),
				Status:   http.StatusBadRequest,
				Cause:    err.Error(),
			}
		}
	}
	c.Set("maxOverall",maxOverall)
	c.Set("minOverall",minOverall)
	return nil
}

func (tc *PlayerController) GetFilter(c *gin.Context) {
	err := tc.validateFilter(c)
	if err != nil {
		c.JSON(http.StatusBadRequest,err)
	}

	minOverall, _ := c.Get("minOverall")
	maxOverall, _ := c.Get("maxOverall")

	players, err := tc.Service.GetPlayers(c, minOverall.(int), maxOverall.(int))
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
func (tc *PlayerController) Load(c *gin.Context) {

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
