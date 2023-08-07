package app

import (
	"flights-api/controllers"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()

	router.GET("/flights/search", controllers.SearchFlights)

	_ = router.Run(":8080")
}
