Group Members: 
Ezekiel Klenicki,
Gabriel Peer-Drake,
Kate Schwitz,
Walter Stahll

"Free Time": A Social Calendar App

Free Time allows people to collaboratively create a schedule to determine the best "free time" for meetings and gatherings. 

# Backend Setup
## Installing Packages

In the event that main.go will not run because of the packages listed in the import section, delete "go.sum" and run the following commands in terminal while in the project directory:

```
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson
go get github.com/joho/godotenv
```

## Gaining server credentials

To prevent anyone from forking the repo and gaining access to the server a .env file has been seperated and shared with the emails of contributors and can be accessed from this link:

```
https://drive.google.com/file/d/1gr3HSneNz5pxcidbdzGAvTGrjfwwbZTf/view?usp=sharing
```

Download and unzip the folder and place the .env file in the same folder as main.go. If the .env file is not visible, look online on how to view hidden files on your respective os (mac, windows10).

## Whitelisting IP Address

Send me (Gabe) your IP address or attempt to run main.go and form a connection with the database. From there I (Gabe) can manually whitelist an IP address when necessary.

## Running the application

navigate to the folder that contains main.go in your terminal and run the snippet:

```
go run main.go
```
You should recieve confirmation that a connection to the database has been made and that a server has successfully launched on localhost:8080

# Frontend Setup
## Installed packages:
Angular:
npm install @angular/core
npm i @angular/platform-browser

Styling:
npm install ngx-bootstrap --save

FullCalendar:
npm i fullcalendar
npm install @fullcalendar/interaction

## Running the application:
opening the terminal and typing ng serve will provide a local host link on which the webpage is displayed (normally "** Angular Live Development Server is listening on localhost:4200, open your browser on http://localhost:4200/ **")

Additionally, to test the Login component, make sure you have the Backend running through a server hosted on localhost:8080. The Login component sends the user's form input to the backend as a JSON file to verify the username and password. This feature will only work if both are running at once.
