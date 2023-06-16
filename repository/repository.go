package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/drewcook/go-mongo-http-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB_HOST = "mongodb://localhost:27017"
const DB_NAME = "go-mongo"

// Connect to a new mongo db instance
func ConnectToMongoDB() *mongo.Database {
	// Create a non-nil context
	ctx := context.TODO()
	// Connect to mongo db client
	clientOptions := options.Client().ApplyURI(DB_HOST)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	// Return the database instance
	return client.Database(DB_NAME)
}

// Get the collection from the mongo db instance
func GetCollection(collectionName string) *mongo.Collection {
	return ConnectToMongoDB().Collection(collectionName)
}

// Seed an array of 10 primitive.M records into the collection
func Seed(collection *mongo.Collection) {
	ctx := context.Background()
	// Create 10 Book records with mock data
	var books []interface{}
	for i := 0; i < 10; i++ {
		book := models.Book{
			Title:  fmt.Sprintf("Book %d Title", i),
			Author: fmt.Sprintf("Book %d Author", i),
		}
		books = append(books, book)
	}
	// Insert the records into the collection
	if _, err := collection.InsertMany(ctx, books); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Seeded the collection!")
}

// Write a handler that GETs all records of books from the collection
func GetBooks(collection *mongo.Collection) []primitive.M {
	// Get context
	ctx := context.Background()
	// Get all records from collection
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	// Transform data into slice of primitive.M
	var data []primitive.M
	for cur.Next(ctx) {
		var record bson.M
		if err := cur.Decode(&record); err == nil {
			log.Fatal(err)
		}
		data = append(data, record)
	}
	return data
}

// func GetAllRecords(collection *mongo.Collection) []primitive.M {
// 	// Get context
// 	ctx := context.Background()
// 	// Get all records from collection
// 	cur, err := collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cur.Close(ctx)
// 	// Transform data into slice of primitive.M
// 	var data []primitive.M
// 	for cur.Next(ctx) {
// 		var record bson.M
// 		if err := cur.Decode(&record); err == nil {
// 			log.Fatal(err)
// 		}
// 		data = append(data, record)
// 	}
// 	return data
// }
