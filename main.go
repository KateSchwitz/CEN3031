package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Printf("Connected to database")
	usersCollection := client.Database("testing").Collection("users")

	// insert a single document into a collection
	// create a bson.D object
	user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), user)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)

	// insert multiple documents into a collection
	// create a slice of bson.D objects
	users := []interface{}{
		bson.D{{"fullName", "User 2"}, {"age", 25}},
		bson.D{{"fullName", "User 3"}, {"age", 20}},
		bson.D{{"fullName", "User 4"}, {"age", 28}},
	}
	// insert the bson object slice using InsertMany()
	results, err := usersCollection.InsertMany(context.TODO(), users)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
