package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestLoginHandler(t *testing.T) {
	reqBody := []byte(`{"username":"Gabe","password":"Password123?"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(loginHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, rr.Code)
	}

	expectedResponseBody := "You have successfully logged in"
	if rr.Body.String() != expectedResponseBody {
		t.Errorf("expected body '%s' but got '%s'", expectedResponseBody, rr.Body.String())
	}
}

func TestAboutEndpoint(t *testing.T) {
	// Create a new HTTP request to the about endpoint
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	//create a mock session

	session, _ := store.Get(req, "session")
	session.Values["user_id"] = "123"
	session.Save(req, rr)

	// Call the aboutHandler function with the HTTP request and response recorder
	handler := http.HandlerFunc(requireLogin(indexHandler))
	handler.ServeHTTP(rr, req)

	// Check that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	session.Options.MaxAge = -1
	err = session.Save(req, rr)
	if err != nil {
		panic(err)
	}

	// Check that the response body contains the expected information about the application

}

func TestRegisterHandler(t *testing.T) {
	// create a new router instance
	r := mux.NewRouter()

	// register the handler function
	r.HandleFunc("/register", registerHandler).Methods("POST")

	// create a new test server with the router
	ts := httptest.NewServer(r)
	defer ts.Close()

	// create a JSON payload with the user's desired username and password
	payload := []byte(`{"username": "Testuser4", "password": "Testpass123!"}`)

	// create a new POST request with the payload
	req, err := http.NewRequest("POST", ts.URL+"/register", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// set the request header to indicate that the payload is JSON
	req.Header.Set("Content-Type", "application/json")

	// create a new HTTP client and send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// check the response status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}
}

func TestAddEventRaw(t *testing.T) {
	reqBody := []byte(`{"title":"Unit Test","color":"Blue","start_date":"02.01.02","end_date":"02.01.02"}`)
	req, err := http.NewRequest("POST", "/addEventRaw", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(eventHandlerRaw)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, rr.Code)
	}

}

func TestDeleteEvent(t *testing.T) {
	reqBody := []byte(`{"title":"Unit Test","color":"Blue","start_date":"02.01.02","end_date":"02.01.02"}`)
	req, err := http.NewRequest("DELETE", "/deleteEvent", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(deleteEventHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("expected status %d but got %d", http.StatusNoContent, rr.Code)
	}

}

func TestEditEvent(t *testing.T) {
	reqBody := []byte(`{"title":"Unit Test","color":"Blue","start_date":"02.01.02","end_date":"02.01.02"}`)
	req, err := http.NewRequest("PUT", "/editEvent", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(editEventHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, rr.Code)
	}

}

func TestDeleteUser(t *testing.T) {
	reqBody := []byte(`{"username":"Testuser3","password":"Testpass123!"}`)
	req, err := http.NewRequest("DELETE", "/deleteUser", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	session, _ := store.Get(req, "session")
	session.Values["user_id"] = "123"
	session.Save(req, rr)

	handler := http.HandlerFunc(requireLogin(deleteAccountHandler))

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("expected status %d but got %d", http.StatusNoContent, rr.Code)
	}

}
