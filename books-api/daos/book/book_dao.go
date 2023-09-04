package book

import (
	model "books-api/models"
	"books-api/utils/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetById(id string) model.Book {
	var book model.Book
	db := db.MongoDb
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return book
	}
	err = db.Collection("books").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&book)
	if err != nil {
		fmt.Println(err)
		return book
	}
	return book

}

func Insert(book model.Book) model.Book {
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
