package services

import (
	bookDao "github.com/IgancioRey/books_microservice/daos/book"
	"github.com/IgancioRey/books_microservice/dtos"
	"github.com/IgancioRey/books_microservice/models"
	e "github.com/IgancioRey/books_microservice/utils/errors"
)

type bookService struct{}
type bookServiceInterface interface {
	InsertBook(bookDto dtos.BookDto) (dtos.BookDto, e.ApiError)
}

var (
	BookService bookServiceInterface
)

func init() {
	BookService = &bookService{}
}

func (s *bookService) InsertBook(bookDto dtos.BookDto) (dtos.BookDto, e.ApiError) {

	var book models.Book

	book.Name = bookDto.Name

	book = bookDao.Insert(book)

	if book.Id.Hex() == "000000000000000000000000" {
		return bookDto, e.NewBadRequestApiError("error in insert")
	}
	bookDto.Id = book.Id.Hex()

	return bookDto, nil
}
