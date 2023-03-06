package services

import (
	"errors"

	bookDao "github.com/IgancioRey/books_microservice/daos/book"
	dtos "github.com/IgancioRey/books_microservice/dtos/book"
	"github.com/IgancioRey/books_microservice/models"
)

type bookService struct{}
type bookServiceInterface interface {
	InsertBook(bookDto dtos.BookDto) (dtos.BookDto, error)
}

var (
	BookService bookServiceInterface
)

func init() {
	BookService = &bookService{}
}

func (s *bookService) InsertBook(bookDto dtos.BookDto) (dtos.BookDto, error) {

	var book models.Book

	book.Name = bookDto.Name

	book = bookDao.Insert(book)

	if book.Id.Hex() == "000000000000000000000000" {
		return bookDto, errors.New("error in insert")
	}
	bookDto.Id = book.Id.Hex()

	return bookDto, nil
}
