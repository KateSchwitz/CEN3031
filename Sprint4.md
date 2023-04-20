# Video Demo:

# FRONTEND
## Accomplished:
  - Proper page routing for the calendar integration
  - Cleaned page navigation
  - Fixed bug in sprint 3 regarding calendar grid displaying over other components
  - CSS Changes
  - Set-up Frontend/Backend connection via login page
  - Improved on Event Adding function
  
## Testing:
  - Unit tests are found in the .spec.ts of each _.component file
  - Cypress tests are found in the .cy.ts of each _.component file

# BACKEND
## Accomplished:
- Added functionality to delete user accounts by DELETE request
- Added functionality to edit an event by PUT request
- Completed unit tests to prove functionality of previous features
- Added functionality to unit tests that include sessions e.g. requireLogin wrapper works properly

# API Documentation
The following endpoints are available in addition to those presented in Sprint2.md

# NEW TO SPRINT 4

## 'DELETE /deleteUser'
This endpoint is used to delete a user from the database, requires a user session to be executed, otherwise redirects to login page.

### Request Parameters
|Parameter|type|Required|Description|
|---|---|---|---|
|'username'|string|yes|username|
|'password'|string|yes|user password|

### Response
If the deletion attempt is successful, the server responds with a status code of '204 NO CONTENT'.

If the deletion attempt fails due to the event not existing in the database, the server responds with a status code of '404 NOT FOUND'

## 'PUT /editEvent'
This endpoint is used to edit a pre existing event. Requires the new event be included in the post request, the new event is filtered to find and replace the old one.

### Request Parameters
|Parameter|type|Required|Description|
|---|---|---|---|
|'title'|string|yes|The name of the event|
|'color'|string|yes|The color of the displayed event ("#000000)|
|'start_date'|string|yes|The start date of the event ("YYYY-MM-DDTHH:MM")|
|'end_date'|string|yes|The end date of the event ("YYYY-MM-DDTHH:MM")|

### Response
If the update is successful, the server responds with status code of '200 OK'. Otherwise, responds with status code '404 NOT FOUND'

# FROM SPRINT 3 AND PRIOR

## 'DELETE /deleteEvent'
This endpoint is used to delete an event from the database collection using filtered components.

### Request Parameters
|Parameter|type|Required|Description|
|---|---|---|---|
|'title'|string|yes|The name of the event|
|'color'|string|yes|The color of the displayed event ("#000000)|
|'start_date'|string|yes|The start date of the event ("YYYY-MM-DDTHH:MM")|
|'end_date'|string|yes|The end date of the event ("YYYY-MM-DDTHH:MM")|

### Response
If the deletion attempt is successful, the server responds with a status code of '204 NO CONTENT'.

If the deletion attempt fails due to the event not existing in the database, the server responds with a status code of '404 NOT FOUND'

## 'POST /addEvent'
This endpoint is used to add an event to the database collection using a POST request with a form-urlencoded body.

### Request Parameters
|Parameter|type|Required|Description|
|---|---|---|---|
|'title'|string|yes|The name of the event|
|'color'|string|yes|The color of the displayed event ("#000000)|
|'start_date'|string|yes|The start date of the event ("YYYY-MM-DDTHH:MM")|
|'end_date'|string|yes|The end date of the event ("YYYY-MM-DDTHH:MM")|

### Response
If the insertion attempt is successful, the server responds with a status code of '200 OK'.

## 'POST /addEventRaw'
This endpoint is used to add an event to the database collection using a POST request with a raw body.

### Request Parameters
|Parameter|type|Required|Description|
|---|---|---|---|
|'title'|string|yes|The name of the event|
|'color'|string|yes|The color of the displayed event ("#000000)|
|'start_date'|string|yes|The start date of the event ("YYYY-MM-DDTHH:MM")|
|'end_date'|string|yes|The end date of the event ("YYYY-MM-DDTHH:MM")|

### Response
If the insertion attempt is successful, the server responds with a status code of '200 OK'.

## 'GET /clearEvents'
This endpoint is used to clear the database collection used for events. Its intended use is for testing purposes.

### Response
If the collection is successfully cleared, the server responds with a status code of '200 OK'.

# Backend Tests
Backend tests can be found in unit_tests.go in the main folder

## Currently supported:
- /login
- /register
- /addEventRaw
- /deleteEvent
- /deleteUser
- /editEvent
- /
Tests see if expected output is returned based on input e.g. login should return http.StatusOK.
