# BACKEND:

## Accomplished:
- Rectified issues that were preventing unit tests from functioning correctly from sprint 2
- Added additional functionality in event components, namely the insertion and deletion of events into the database
- Created additional unit tests to verify functionality of new features
## In progress:
- Incorporate session functionality into event handling
- Standardization of document-style database
- Determine and impliment routing between pages
## To Do for Sprint 4:
- Ensure that a user may only access events that they have created
- Figure out and impliment event retrieval from the database in a way that is displayable to the frontend

# API Documentation
The following endpoints are available in addition to those presented in Sprint2.md

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

# Backend Tests
Backend tests can be found in unit_tests.go in the main folder

## Currently supported:
- /login
- /register
- /addEventRaw
- /deleteEvent

Tests see if expected output is returned based on input e.g. login should return http.StatusOK.
Currently tests are now completely working because the issue with mongodb has been resolved
