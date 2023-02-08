package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"unicode"

	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var tpl *template.Template

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register Handler Running")
	tpl.ExecuteTemplate(w, "register.html", nil)
}

func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RegisterAuthHandler Running")

	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println("username: ", username)
	fmt.Println("password: ", password)

	var alphaNumeric = true
	for _, char := range username {
		if unicode.IsNumber(char) == false && unicode.IsLetter(char) == false {
			alphaNumeric = false
		}
	}

	var usernameLength bool
	if 4 <= len(username) && 16 >= len(username) {
		usernameLength = true
	}
	println("alphaNumeric: ", alphaNumeric)
	println("usernameLength: ", usernameLength)

	var pswdLowerCase, pswdUpperCase, pswdNumber, pswdSpecialChar, pswdLength, pswdNoSpace bool
	pswdNoSpace = true

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			pswdLowerCase = true

		case unicode.IsUpper(char):
			pswdUpperCase = true
		case unicode.IsNumber(char):
			pswdNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecialChar = true
		case unicode.IsSpace(int32(char)):
			pswdNoSpace = false
		}
	}

	if 8 <= len(password) && 60 >= len(password) {
		pswdLength = true
	}

	fmt.Println("pswdLowerCase: ", pswdLowerCase, "\npswdUpperCase: ", pswdUpperCase, "\npswdNumber: ", pswdNumber, "\npswdSpecialChar: ", pswdSpecialChar, "\npswdLength: ", pswdLength, "\npswdNoSpace: ", pswdNoSpace)
	if !pswdLowerCase || !pswdUpperCase || !pswdNumber || !pswdSpecialChar || !pswdLength || !pswdNoSpace {
		tpl.ExecuteTemplate(w, "register.html", "please check username and password criteria")
		return
	}

	// TODO: check if username exists in db already

	var hash []byte
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
		return
	}
	fmt.Println("hash: ", hash)
	fmt.Println("string(hash): ", string(hash))

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	usersCollection := client.Database("testing").Collection("register")

	user := bson.D{{"username", username}, {"password", string(hash)}}

	result, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(nil)
	}
	fmt.Println(result.InsertedID)

}

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
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerAuth", registerAuthHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
