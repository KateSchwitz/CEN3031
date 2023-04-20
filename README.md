Group Members: 
Ezekiel Klenicki,
Gabriel Peer-Drake,
Kate Schwitz,
Walter Stahll

"Free Time": A Social Calendar App

Free Time allows people to collaboratively create a schedule to determine the best "free time" for meetings and gatherings. 

# Backend Setup
This project was written and run using VS Code's Go extension.
## Installing Packages

In the event that the program will not run because of the packages listed in the import section, import the following packages in your terminal while in the project directory. i.e. run: go get github.com/gorilla/context for each dependency listed.

```
"github.com/gorilla/context"
"github.com/gorilla/sessions"
"github.com/joho/godotenv"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
"go.mongodb.org/mongo-driver/mongo/readpref"
"go.mongodb.org/mongo-driver/bson"
"golang.org/x/crypto/bcrypt"

```

## Gaining server credentials

Create a file titled '.env' and paste the following two lines in:

```
export MONGODB_URI = "mongodb+srv://project_group_5:XbDuA0Vid6BnuRY7@cluster0.5ch00jt.mongodb.net/?retryWrites=true&w=majority"
export SESSION_KEY = "av0iweruvggwpervwerifpieu"
```

## Whitelisting IP Address

All IP addresses are whitelisted, meaning access to the database should not be an issue.

## Running the application

navigate to the folder that contains main.go in your terminal and run the command:

```
go run main.go login.go register.go event.go
```
You should recieve confirmation that a connection to the database has been made and that a server has successfully launched on localhost:8080

# Frontend Setup
## Installed packages:
Angular:
• npm install -g @angular/cli
• npm install @angular/core
• npm i @angular/platform-browser

Styling:
• npm install ngx-bootstrap --save

FullCalendar:
• npm i fullcalendar
• npm install @fullcalendar/interaction

## Running the application:
opening the terminal and typing ng serve will provide a local host link on which the webpage is displayed (normally "** Angular Live Development Server is listening on localhost:4200, open your browser on http://localhost:4200/ **")

Additionally, to test the Login component, make sure you have the Backend running through a server hosted on localhost:8080. The Login component sends the user's form input to the backend as a JSON file to verify the username and password. This feature will only work if both are running at once.
