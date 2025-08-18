package controllers

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/driver/monogocrypt/options"
)

const connectionString string = "mongodb+srv://Pratik:Pratik5722@cluster0.s79v2th.mongodb.net/"
const dbName string = "netflix"
const collectionName string = "watchlist"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions) 
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance is ready")
}

func insertOneMovie(movie models.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in database with ID:", inserted.InsertedID)
}