package router

import (
	"fmt"
	bookController "github.com/IgancioRey/books_microservice/controllers/book"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.POST("/books", bookController.Insert)
	router.GET("/books", bookController.GetBooks)
	router.GET("/books/:id", bookController.Get)
	/*
		router.GET("/books", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	*/

	fmt.Println("mappings configurations ok!")
}
