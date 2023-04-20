package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct {
	EventName   string
	Color       string // "#000000" format
	DTCreate    string // Date & Time that Event was created
	DTStart     string // Date & Time that Event is scheduled to start
	DTEnd       string // Date & Time that Event is scheduled to end
	CreatorName string // Event Creator's ID
	CreatorID   string // Event Creator's ID
	EventDesc   string // Event Description
}

// inserts an event into a database
func insertEvent(collection *mongo.Collection, newEvent Event) {
	insertResult, err := collection.InsertOne(context.TODO(), newEvent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document:", insertResult.InsertedID)
}

// clears all events from the Events collection on the MongoDB database
func clearEvents(w http.ResponseWriter, r *http.Request) {
	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	eventsCollection := client.Database("testing").Collection("events")
	err = eventsCollection.Drop(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Collection dropped successfully!")
}

// Creates a new event in the Events collection Post request with form-urlencoded body
func eventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****eventHandler Running*****")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	color := r.FormValue("color")
	start_date := r.FormValue("start_date")
	end_date := r.FormValue("end_date")
	println(title)
	println(color)
	println(start_date)
	println(end_date)

	ct := time.Now()
	creatorName := "placeholderName"
	creatorID := "placeholderID"
	eventDesc := "This is a placeholder for an event description."

	newEvent := Event{EventName: title, Color: color, DTCreate: ct.Format("2006-01-02 15:04"),
		DTStart: start_date, DTEnd: end_date, CreatorName: creatorName, CreatorID: creatorID, EventDesc: eventDesc}

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	eventsCollection := client.Database("testing").Collection("events")

	insertEvent(eventsCollection, newEvent)
}

// Creates a new event in the Events collection Post request with raw body
func eventHandlerRaw(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****eventHandlerRaw Running*****")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		println("method not allowed")
		return
	}

	var event struct {
		Title      string "json:'title'"
		Color      string "json:'color'"
		Start_date string "json:'start_date'"
		End_date   string "json:'end_date'"
	}
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		println("bad request")
		return
	}
	println(event.Title)
	println(event.Color)
	println(event.Start_date)
	println(event.End_date)

	ct := time.Now()
	creatorName := "placeholderName"
	creatorID := "placeholderID"
	eventDesc := "This is a placeholder for an event description."

	newEvent := Event{EventName: event.Title, Color: event.Color, DTCreate: ct.Format("2006-01-02 15:04"),
		DTStart: event.Start_date, DTEnd: event.End_date, CreatorName: creatorName, CreatorID: creatorID, EventDesc: eventDesc}

	uri := "mongodb+srv://project_group_5:XbDuA0Vid6BnuRY7@cluster0.5ch00jt.mongodb.net/?retryWrites=true&w=majority" //os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	eventsCollection := client.Database("testing").Collection("events")

	insertEvent(eventsCollection, newEvent)
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		println("method not allowed")
		return
	}

	var event struct {
		Title      string "json:'title'"
		Color      string "json:'color'"
		Start_date string "json:'start_date'"
		End_date   string "json:'end_date'"
	}

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		println("bad request")
		return
	}

	println(event.Title)
	println(event.Color)
	println(event.Start_date)
	println(event.End_date)

	uri := "mongodb+srv://project_group_5:XbDuA0Vid6BnuRY7@cluster0.5ch00jt.mongodb.net/?retryWrites=true&w=majority" //os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	eventsCollection := client.Database("testing").Collection("events")

	filter := bson.D{{"eventname", event.Title}, {"color", event.Color}, {"dtstart", event.Start_date}, {"dtend", event.End_date}}

	result, err := eventsCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	if result.DeletedCount == 1 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func editEventHandler(w http.ResponseWriter, r *http.Request) {
	// make sure it is a post request
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		println("method not allowed")
		return
	}

	// contents of post request
	var event struct {
		Title      string "json:'title'"
		Color      string "json:'color'"
		Start_date string "json:'start_date'"
		End_date   string "json:'end_date'"
	}
	update := bson.D{{"$set", bson.D{
		{"eventname", event.Title},
		{"color", event.Color},
		{"dtstart", event.Start_date},
		{"dtend", event.End_date},
	}}}

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		println("bad request")
		return
	}

	uri := "mongodb+srv://project_group_5:XbDuA0Vid6BnuRY7@cluster0.5ch00jt.mongodb.net/?retryWrites=true&w=majority" //os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	eventsCollection := client.Database("testing").Collection("events")

	filter1 := bson.D{{"eventname", event.Title}}

	result1, err := eventsCollection.UpdateOne(context.TODO(), filter1, update)

	filter2 := bson.D{{"color", event.Color}}

	result2, err := eventsCollection.UpdateOne(context.TODO(), filter2, update)

	if result1.MatchedCount == 1 || result2.MatchedCount == 1 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
