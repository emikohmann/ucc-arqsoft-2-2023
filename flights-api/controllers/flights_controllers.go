package controllers

import (
	"flights-api/services"
	"flights-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchFlights(c *gin.Context) {

	flights := services.SearchFlights(c)

	c.JSON(http.StatusOK, flights)
}

func GenerateUUID(c *gin.Context) {
	requestID := utils.NewUUID()
	c.Set("request_id", requestID)
}
