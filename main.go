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

	// retrieve username and password from the form
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println("username: ", username)
	fmt.Println("password: ", password)

	// check if username is only alphanumeric characters
	var alphaNumeric = true
	for _, char := range username {
		if unicode.IsNumber(char) == false && unicode.IsLetter(char) == false {
			alphaNumeric = false
		}
	}

	// check if username is appropriate length
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
		case unicode.IsLower(char): // check if password has a lowercase char
			pswdLowerCase = true

		case unicode.IsUpper(char): // check if password has an uppercase char
			pswdUpperCase = true
		case unicode.IsNumber(char): // check if password has a number
			pswdNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char): // check if password has a special character
			pswdSpecialChar = true
		case unicode.IsSpace(int32(char)): // check if password has no spaces
			pswdNoSpace = false
		}
	}

	// check if password is appropriate length
	if 8 <= len(password) && 60 >= len(password) {
		pswdLength = true
	}

	fmt.Println("pswdLowerCase: ", pswdLowerCase, "\npswdUpperCase: ", pswdUpperCase, "\npswdNumber: ", pswdNumber, "\npswdSpecialChar: ", pswdSpecialChar, "\npswdLength: ", pswdLength, "\npswdNoSpace: ", pswdNoSpace)
	if !pswdLowerCase || !pswdUpperCase || !pswdNumber || !pswdSpecialChar || !pswdLength || !pswdNoSpace {
		tpl.ExecuteTemplate(w, "register.html", "please check username and password criteria")
		return
	}

	// check if username exists in db already
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

	if len(results) > 0 {
		tpl.ExecuteTemplate(w, "register.html", "the username you selected is already in use")
		return
	}

	// create a hash for password
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
		return
	}
	fmt.Println("hash: ", hash)
	fmt.Println("string(hash): ", string(hash))

	user := bson.D{{"username", username}, {"password", string(hash)}}

	// insert username and str(hash) into database
	result, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(nil)
	}
	fmt.Println(result.InsertedID)

	fmt.Fprint(w, "Your account has been successfully created")

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
