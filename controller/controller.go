package controller

import (
	"context"
	"fmt"
	"log"

	""
)

// variable for connect form the MongoDB
const connectionString = "mongodb+srv://test-db:test@cluster0.wv5caal.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

//MOST IMPORTANT
var collection *mongo.collection

//Connect with mongoDB

func init() {
	//client option

	clientOption := options.Client().ApplyURL(connectionString)

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
