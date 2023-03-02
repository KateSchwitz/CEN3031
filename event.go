package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"time"

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

func eventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****eventHandler Running*****")
	tpl.ExecuteTemplate(w, "app.component.html", nil)
}

func eventAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****eventAuthHandler Running*****")
	r.ParseForm()

	title := r.FormValue("title")
	color := r.FormValue("color")
	ct := time.Now()
	start := r.FormValue("start_date")
	end := r.FormValue("end_date")
	creatorName := "placeholderName"
	creatorID := "placeholderID"
	eventDesc := "This is a placeholder for an event description."

	newEvent := Event{EventName: title, Color: color, DTCreate: ct.Format("2006-01-02 15:04"),
		DTStart: start, DTEnd: end, CreatorName: creatorName, CreatorID: creatorID, EventDesc: eventDesc}

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	usersCollection := client.Database("testing").Collection("events")

	insertEvent(usersCollection, newEvent)
}
