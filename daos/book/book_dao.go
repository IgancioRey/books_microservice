package book

import (
	"context"
	"fmt"
	"github.com/IgancioRey/books_microservice/models"
	"github.com/IgancioRey/books_microservice/utils/db"
	"go.mongodb.org/mongo-driver/bson"
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

func GetBooks() []models.Book {
	db := db.MongoDb
	cur, err := db.Collection("books").Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var books []models.Book
	for cur.Next(context.TODO()) {
		var book models.Book
		err := cur.Decode(&book)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		books = append(books, book)
	}

	defer cur.Close(context.TODO())

	return books
}

func GetById(id string) models.Book {
	var book models.Book
	db := db.MongoDb
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return book
	}
	err = db.Collection("books").FindOne(context.TODO(), bson.D{{Key: "_id", Value: objID}}).Decode(&book)
	if err != nil {
		fmt.Println(err)
		return book
	}
	return book

}
