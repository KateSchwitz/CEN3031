Group Members: 
Ezekiel Klenicki,
Gabriel Peer-Drake,
Kate Schwitz,
Walter Stahll

"Free Time": A Social Calendar App

Features: personal calendar to link free time with collaborators
Each personal calendar keeps track of the user’s events.
A user is notified when an event is shared with them.
Each event can be public, private, hidden, or manually selected
A public event shows an event’s name and details to everyone on your friends list.
A private event blocks out a period of time where only you can see the designated name and details of the event. Still shows everyone that you’re busy during the time of the event.
A hidden event only appears on a calendar when the event’s creator is the one viewing it. There will be no event shown for anyone else.
Manual selection can be used to go through your Friends List and mark which people can see an event’s name and details (events could be specified to be public, private, and hidden for certain people).
Friends can be added with someone else’s username, so long as they accept your request. Specific accounts should be able to be blocked from sending requests.
There could be a specific event type called “Free Time” that shows up on other people’s calendars (maybe only when the setting below is enabled). These kinds of events could be either public, hidden, or manually selected, but not private. (bc if you wanted private free time, why would you go out of your way to make an event for it?)
Maybe your personal calendar can be set to only show Free Time (yours and others), to make it easier to determine which people you have overlapping Free Time with.

# Backend Setup
## Installing Packages

In the event that main.go will not run because of the packages listed in the import section, delete "go.sum" and run the following commands in terminal while in the project directory:

```
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson
go get github.com/joho/godotenv

go get golang.org/x/crypto/bcrypt
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