package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"time"

	contextG "github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var tpl *template.Template

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type Event struct {
	EventName   string
	DTCreate    string // Date & Time that Event was created
	DTStart     string // Date & Time that Event is scheduled to start
	DTEnd       string // Date & Time that Event is scheduled to end
	CreatorName string // Event Creator's ID
	CreatorID   string // Event Creator's ID
	EventDesc   string // Event Description
}

func insertEvent(collection *mongo.Collection, newEvent Event) {
	insertResult, err := collection.InsertOne(context.TODO(), newEvent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document:", insertResult.InsertedID)
}

func main() {
	var err error
	tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println("Parsing Templates Error:")
		panic(err.Error)
	}

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

	// testing the insertEvent function
	currentTime := time.Now() // gets the current time
	startTime := time.Date(currentTime.Year(), time.March, 11, 0, 0, 0, 0, currentTime.Location())
	endTime := time.Date(currentTime.Year(), time.March, 19, 23, 59, 59, 0, currentTime.Location())
	newEvent := Event{EventName: "Spring Break 2023", DTCreate: currentTime.Format("2006-01-02 15:04:05"), DTStart: startTime.Format("2006-01-02 15:04:05"), DTEnd: endTime.Format("2006-01-02 15:04:05"), CreatorName: "stahllw", CreatorID: "1", EventDesc: "2023 Spring Break! We have no school!"}
	eventsCollection := client.Database("testing").Collection("events")
	err = eventsCollection.Drop(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Collection dropped successfully!")
	insertEvent(eventsCollection, newEvent)

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerAuth", registerAuthHandler)
	http.HandleFunc("/login", loginHander)
	http.HandleFunc("/loginAuth", loginAuthHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/", indexHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", contextG.ClearHandler(http.DefaultServeMux)); err != nil {
		log.Fatal(err)
	}
}
