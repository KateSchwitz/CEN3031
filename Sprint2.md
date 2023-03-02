# FRONTEND:

## Accomplished:
- Set up loginPage component to prepare for page routing
- Created basic Cypress test for loginPage component
      Pass unit test of incrementing/decrementing Event Counter
      Input of username/password and clicking Save button

## In progress:
- Pop-up bootstrap modal upon "Add Event" button click
- Properly set up page routing between calendar and login components

## To Do for Sprint 3: 
- Add events via button
- Navigation between pages

# BACKEND:

## Accomplished:
- Realized that handler function setup performed for sprint 1 was incompatable with frontend.
- Restructured handler function to behave more closeley with REST behaviors
- Began adding session functionality (keeping users logged in, restricting access to non-logged-in users)
## In progress:
- Add event functionality to api
- Standardizing and implimenting error codes in an effective way
- Determine and impliment routing between pages
## To Do for Sprint 3:
- Complete the integration of backend and frontend (95% complete)
- Lock down event addition and deletion into the database from frontend

# Progress videos:
https://youtu.be/ulztURZTP-I
https://youtu.be/meY56Wi_7FA

# API Documentation
The following endpoints are available:

## 'POST /register'
This endpoint is used to register a user with a username and password

### Request Parameters
|Parameter|type|Required|Description|
|---|---|---|---|
|'username'|string|yes|The username of the user to register|
|'password'|string|yes|The password of the user to register|

### Response
If the register attempt is successful, the server responds with a status code of '200 OK'.

If the register attempt fails, due to username already being registered the server responds with a status code of '409 Conflict' and returns an error message in the response body.

If the register attempt fails, due to any other reason the server responds with a status code of '401 Unauthorized' and returns an error message in the response body.

## 'POST /login'
This endpoint is used to login a user with a username and password

### Request Parameters
|Parameter|type|Required|Description|
|---|---|---|---|
|'username'|string|yes|The username of the user to log in|
|'password'|string|yes|The password of the user to log in|

### Response
If the log in attempt is successful, the server responds with a status code of '200 OK' and sets a session cookie for the user.

If the log in attempt fails the server responds with a status code of '401 Unauthorized' and returns an error message in the response body.

## 'GET /about'
This endpoint is used to retrieve information about the application.

### Response
If the user is logged in, the server responds with a status code of 200 OK and returns information about the application in the response body.

If the user is not logged in, the server redirects them to the login page.

## 'POST /logout'
This endpoint is used to log out the currently logged in user.

### Response
If the logout attempt is successful, the server responds with a status code of 302 Found and redirects the user to the login page.

## Example
Here's an example of how you can use the API to log in a user:

```
POST /login HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "username": "johndoe",
  "password": "password123"
}
```
And here's an example of a successful response:
```
HTTP/1.1 200 OK
Set-Cookie: session-key=abc123; Path=/; HttpOnly
```
# Backend Tests
Backend tests can be found in unit_tests.go in the main folder

## Currently supported:
- /login
- /register
- /about

Tests see if expected output is returned based on input e.g. login should return http.StatusOK.
Currently tests are not completely working because of an issue with mongodb, however they are able to recieve input and use it, which should indicate that the body of loginHandler, registerHandler and aboutHandler are otherwise working. Postman testing confirms this.
