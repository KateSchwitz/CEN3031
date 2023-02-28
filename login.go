package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****indexHandler running*****")
	session, _ := store.Get(r, "session")
	_, ok := session.Values["userID"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound) // http.StatusFound is 302
		return
	}
	tpl.ExecuteTemplate(w, "index.html", "Logged In")
}

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

		session, _ := store.Get(r, "session")
		session.Values["userID"] = username
		session.Save(r, w)

		tpl.ExecuteTemplate(w, "index.html", "Logged In")
		return
	}

	// correct username and incorrect password sends back to login
	fmt.Println("Incorrect password")
	tpl.ExecuteTemplate(w, "login.html", "Please check username and password")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****logoutHandler running*****")
	session, _ := store.Get(r, "session")

	delete(session.Values, "username")
	session.Save(r, w)
	tpl.ExecuteTemplate(w, "login.html", "logged out")
}
