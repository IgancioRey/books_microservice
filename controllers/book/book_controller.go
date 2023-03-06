package book

import (
	"fmt"
	"net/http"

	dtos "github.com/IgancioRey/books_microservice/dtos/book"
	service "github.com/IgancioRey/books_microservice/services/book"

	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	var bookDto dtos.BookDto
	err := c.BindJSON(&bookDto)

	// Error Parsing json param
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	bookDto, er := service.BookService.InsertBook(bookDto)

	// Error del Insert
	if er != nil {
		c.JSON(http.StatusInternalServerError, er)
		return
	}

	c.JSON(http.StatusCreated, bookDto)
}
