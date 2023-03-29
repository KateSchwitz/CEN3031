package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"unicode"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****RegisterAuthHandler Running*****")

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
	println(credentials.Username)
	println(credentials.Password)

	// check if username is only alphanumeric characters
	var alphaNumeric = true
	for _, char := range credentials.Username {
		if unicode.IsNumber(char) == false && unicode.IsLetter(char) == false {
			alphaNumeric = false
		}
	}

	// check if username is appropriate length
	var usernameLength bool
	if 4 <= len(credentials.Username) && 16 >= len(credentials.Username) {
		usernameLength = true
	}
	println("alphaNumeric: ", alphaNumeric)
	println("usernameLength: ", usernameLength)

	var pswdLowerCase, pswdUpperCase, pswdNumber, pswdSpecialChar, pswdLength, pswdNoSpace bool
	pswdNoSpace = true

	for _, char := range credentials.Password {
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
	if 8 <= len(credentials.Password) && 60 >= len(credentials.Password) {
		pswdLength = true
	}

	fmt.Println("pswdLowerCase: ", pswdLowerCase, "\npswdUpperCase: ", pswdUpperCase, "\npswdNumber: ", pswdNumber, "\npswdSpecialChar: ", pswdSpecialChar, "\npswdLength: ", pswdLength, "\npswdNoSpace: ", pswdNoSpace)
	if !pswdLowerCase || !pswdUpperCase || !pswdNumber || !pswdSpecialChar || !pswdLength || !pswdNoSpace {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// check if username exists in db already
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

	if len(results) > 0 {
		http.Error(w, "Selected username already in use", http.StatusConflict)

		return
	}

	// create a hash for password
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println("hash: ", hash)
	fmt.Println("string(hash): ", string(hash))

	user := bson.D{{"username", credentials.Username}, {"password", string(hash)}}

	// insert username and str(hash) into database
	result, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(nil)
	}
	fmt.Println(result.InsertedID)

	fmt.Fprint(w, "Your account has been successfully created")

}
