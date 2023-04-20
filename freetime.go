package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FreeTime struct {
	DTStart     string
	DTEnd       string
	CreatorName string
	CreatorID   string
}

// clears all events from the Events collection on the MongoDB database
func clearFreeTime(w http.ResponseWriter, r *http.Request) {
	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	freeTimeCollection := client.Database("testing").Collection("freeTime")
	err = freeTimeCollection.Drop(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Collection dropped successfully!")
}

// inserts freetime into a database
func insertFreeTime(freeTimeCollection *mongo.Collection, newFreeTime FreeTime) {
	insertResult, err := freeTimeCollection.InsertOne(context.TODO(), newFreeTime)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted free time from the dates:", newFreeTime.DTStart, "-", newFreeTime.DTEnd)
	fmt.Println("Free Time ID: ", insertResult.InsertedID)
}

/*
// turns the time in a day that is not being used by another event into free time
func freeTimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****freeTimeHandler Running*****")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		println("method not allowed")
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dtStart := r.FormValue("start_date")
	dtEnd := r.FormValue("end_date")
	creatorName := "placeholderName"
	creatorID := "placeholderID"

	newFreeTime := FreeTime{DTStart: dtStart, DTEnd: dtEnd, CreatorName: creatorName, CreatorID: creatorID}

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	freeTimeCollection := client.Database("testing").Collection("freeTime")
	insertFreeTime(freeTimeCollection, newFreeTime)
}
*/

func freeTimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		println("method not allowed")
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dtStart := r.FormValue("start_date")
	dtEnd := r.FormValue("end_date")
	creatorName := "placeholderName"
	creatorID := "placeholderID"

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	eventsCollection := client.Database("testing").Collection("events")
	freeTimeCollection := client.Database("testing").Collection("freeTime")

	filter := bson.M{
		"dtstart": bson.M{"$lte": dtStart},
		"dtend":   bson.M{"$gte": dtEnd},
	}
	var event Event
	err = eventsCollection.FindOne(context.Background(), filter).Decode(&event)
	fmt.Println(dtStart)
	fmt.Println(dtEnd)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			newFreeTime := FreeTime{DTStart: dtStart, DTEnd: dtEnd, CreatorName: creatorName, CreatorID: creatorID}
			insertFreeTime(freeTimeCollection, newFreeTime)
			fmt.Println("FreeTime inserted normally.")
			return
		}
		return
	}

	freeTimeBefore := FreeTime{DTStart: dtStart, DTEnd: event.DTStart, CreatorName: creatorName, CreatorID: creatorID}
	freeTimeAfter := FreeTime{DTStart: event.DTEnd, DTEnd: dtEnd, CreatorName: creatorName, CreatorID: creatorID}
	insertFreeTime(freeTimeCollection, freeTimeBefore)
	insertFreeTime(freeTimeCollection, freeTimeAfter)
}
