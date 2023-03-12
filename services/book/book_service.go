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
	GetBooks() ([]dtos.BookDto, error)
	GetBook(id string) (dtos.BookDto, error)
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

func (s *bookService) GetBooks() ([]dtos.BookDto, error) {

	booksResult := make([]dtos.BookDto, 0)
	books := bookDao.GetBooks()
	var book dtos.BookDto
	for _, b := range books {
		book.Id = b.Id.Hex()
		book.Name = b.Name

		booksResult = append(booksResult, book)
	}

	return booksResult, nil

}

func (s *bookService) GetBook(id string) (dtos.BookDto, error) {

	var book models.Book = bookDao.GetById(id)
	var bookDto dtos.BookDto

	if book.Id.Hex() == "000000000000000000000000" {
		return bookDto, errors.New("book not found")
	}
	bookDto.Name = book.Name
	bookDto.Id = book.Id.Hex()
	return bookDto, nil
}
