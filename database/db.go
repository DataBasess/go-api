package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mongo() *mongo.Client {
	var uri string
	uri = "mongodb+srv://anusondd:kOEh9rLXvnsXZyIh@testanusorn-fbne4.mongodb.net"
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Connected to MongoDB Fail")
	}

	// fmt.Println("Connected to MongoDB!")
	return client
}
