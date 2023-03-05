package router

import (
	"fmt"
	bookController "github.com/IgancioRey/books_microservice/controllers/book"
	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.POST("/books", bookController.Insert)

	fmt.Println("mappings configurations ok!")
}
