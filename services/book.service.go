package services

import (
	"context"
	"fmt"
	"go-api/database"
	"go-api/models"
	"go-api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

var bookCollection = database.Mongo().Database("goapi").Collection("books")

func GetBooks() (interface{}, error) {
	filter := bson.D{{}}
	cur, findError := bookCollection.Find(context.TODO(), filter)
	if findError != nil {
		fmt.Println(findError)
		return cur, findError
	}
	result := utils.DocumentList(cur)
	return result, findError
}

func GetBook(ID string) (interface{}, error) {
	oid, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": oid}

	var doc models.Book
	err := bookCollection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return doc, nil
}

func InsertBook(doc models.Book) (interface{}, error) {
	result, err := bookCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result.InsertedID, err
}

func UpdateBook(id string, doc models.Book) (interface{}, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}

	mapDoc, err := utils.StructToMap(doc)
	fmt.Println(err, mapDoc)

	update := bson.M{
		"$set": mapDoc,
	}
	result, err := bookCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result.UpsertedID, err
}

func DeleteBook(ID string) (interface{}, error) {
	oid, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": oid}
	result, err := bookCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result.DeletedCount, err
}
