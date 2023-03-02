package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

/*
func loginHander(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****loginHandler Running*****")
	tpl.ExecuteTemplate(w, "login.html", nil)
}
func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****loginAuthHandler Running*****")
	r.ParseForm()

	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println("Username: ", username)
	fmt.Println("Password: ", password)

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	usersCollection := client.Database("testing").Collection("register")

	filter := bson.D{{"username", username}}
	cursor, err := usersCollection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// if username is not registered send back to login
	if len(results) == 0 {
		tpl.ExecuteTemplate(w, "login.html", "Please check username and password")
		return
	}
	var hash string = results[0]["password"].(string)

	// compare password with hash in db
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		fmt.Fprint(w, "You have successfully logged in")
		return
	}

	// correct username and incorrect password sends back to login
	fmt.Println("Incorrect password")
	tpl.ExecuteTemplate(w, "login.html", "Please check username and password")
}

*/

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****loginAuthHandler Running*****")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Username string "json:'username'"
		Password string "json:'password'"
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Username: ", credentials.Username)
	fmt.Println("Password: ", credentials.Password)

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	usersCollection := client.Database("testing").Collection("register")

	filter := bson.D{{"username", credentials.Username}}
	cursor, err := usersCollection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// if username is not registered send back to login
	if len(results) == 0 {
		tpl.ExecuteTemplate(w, "login.html", "Please check username and password")
		return
	}
	var hash string = results[0]["password"].(string)

	// compare password with hash in db
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(credentials.Password))
	if err == nil {
		fmt.Fprint(w, "You have successfully logged in")
		return
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
