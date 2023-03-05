package book

import (
	"fmt"
	"github.com/IgancioRey/books_microservice/dtos"
	service "github.com/IgancioRey/books_microservice/services"
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, bookDto)
}
