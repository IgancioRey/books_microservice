package book

import (
	"context"
	"fmt"

	"github.com/IgancioRey/books_microservice/models"
	"github.com/IgancioRey/books_microservice/utils/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(book models.Book) models.Book {
	db := db.MongoDb
	insertBook := book
	insertBook.Id = primitive.NewObjectID()
	_, err := db.Collection("books").InsertOne(context.TODO(), &insertBook)

	if err != nil {
		fmt.Println(err)
		return book
	}
	book.Id = insertBook.Id
	return book
}
