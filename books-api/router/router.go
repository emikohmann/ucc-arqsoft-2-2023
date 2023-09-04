package router

import (
	bookController "books-api/controllers/book"
	"fmt"
	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.GET("/books/:id", bookController.Get)
	router.POST("/books", bookController.Insert)

	fmt.Println("Finishing mappings configurations")
}
