package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/souravsk/GO-MongoDB-API/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// variable for connect form the MongoDB
const connectionString = "mongodb+srv://test-db:test@cluster0.wv5caal.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

//MOST IMPORTANT
var collection *mongo.Collection

//Connect with mongoDB

func init() {
	//client option

	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success")

	//this is for connecting inside the database to get the collection in the database
	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

//MONGODB helpers - file

//insert 1 record

func insertOneMovie(movie model.Nitflix) { //model is from pacakge model and Netfile is what we create in model file as struct or classs
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted One Movie in database with ID", inserted.InsertedID)
}

//Update One Movie or record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)       //we used this primitive.ObjectIDFromHex so that it is accssepted my mongodb
	filter := bson.M{"_id": id}                       //we are find the id in mongodb it can find aoutomatic
	update := bson.M{"$set": bson.M{"watched": true}} // we are not checking the data we just update the data

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count", result.ModifiedCount)
}

//delete one Movie or record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie get deleted with deleted count:", deleteCount)
}

//delete all records from monogdb

func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movie deletd is ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//get all the movies

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

//Actual Controller -file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}

//create movies

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Nitflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

//mark as watched

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete on movie

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

//Delete all movies

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	countMovie := deleteAllMovie()
	json.NewEncoder(w).Encode(countMovie)
}
