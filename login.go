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

func requireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session")

		if err != nil || !session.IsNew {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****indexHandler running*****")

	info := map[string]string{
		"name":    "My Application",
		"version": "1.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

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
	if err != nil {
		panic(err)
	}

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

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	var hash string = results[0]["password"].(string)

	// compare password with hash in db
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(credentials.Password))
	if err == nil {

		session, _ := store.Get(r, "session")
		session.Values["userID"] = credentials.Username
		session.Save(r, w)

		fmt.Fprint(w, "You have successfully logged in")
		return

	} else {

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****logoutHandler running*****")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method - not post", http.StatusMethodNotAllowed)
		return
	}

	session, _ := store.Get(r, "session")

	delete(session.Values, "userID")
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
}
