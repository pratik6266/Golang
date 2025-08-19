package controllers

import (
	"context"
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"web/models"
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

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Modified %v movies in the database\n", result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v movie from the database\n", result.DeletedCount)
}

func deleteAllMovies(){
	filter := bson.M{}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v movies from the database\n", result.DeletedCount)
}

func getAllMovies() []bson.M {
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var movies []bson.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}


//! controllers
func GetAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
	return
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
	return 
}

func MarkMovieAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	movieId := params["id"]
	updateOneMovie(movieId)
	json.NewEncoder(w).Encode("Movie marked as watched")
	return
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	movieId := params["id"]
	deleteOneMovie(movieId)
	json.NewEncoder(w).Encode("Movie deleted successfully")
	return
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllMovies()
	json.NewEncoder(w).Encode("All movies deleted successfully")
	return
}