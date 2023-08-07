package controllers

import (
	"flights-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchFlights(c *gin.Context) {
	flights := services.SearchFlights(c)

	c.JSON(http.StatusOK, flights)
}

func GenerateUUID(c *gin.Context) {

}
