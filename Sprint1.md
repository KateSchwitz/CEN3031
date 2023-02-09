USER STORIES:

As a busy college student, I want to plan things without having a back and forth conversation about what times work or don’t, so I can enjoy time with my friends

As a busy person, I want to be able to plan things with friends without offering different times that work for both of us so that we can make plans faster.

As a busy college student, I want to be able to showcase what events I have planned throughout my week so that people know when I am free and ask me for plans in those times.

As a social person, I want to see what  events my friends are going to so I can go with them

As an antisocial person, I want to see if my friends are going to the same events I am going to so I do not go alone

As a club leader, I want to be able to see how many people in my club are going to my events, so that I can have an idea of how to accommodate my plans for that group of people.

As a club member, I want to see how many other people are going to club events so that I know I am not the only one going.

As a friend, I want to see when my other friends are free so I can ask them for plans during times I know they won’t be busy

As parents, we want to see when our children are free so that we can plan events to spend time with then

As a rising professional, I want to organize my time with others, so I can be more productive.




Required Pages:
Personal calendar
	Week wise basis
	Interface for adding events (public/private/hidden, event tags/categories)
	Repeating events
Calendar group, only with “free time” to demonstrate overlapping freetime
	Individuals have lighter opacity
	Group overlap has solid color
	Ability to click on a certain member’s name and see their “free time”
Account handling
	Adding friends
	Sharing calendar to group (link via text?)
	Profile page/account info
	“Inbox” for notifications (rsvp to shared events, etc)
	

Soft Timeline:
Sprint 1: 
	.ics file?
	Backend: setup database, with input from form
	Frontend: display a week on a calendar, possibly adding an event

Front End:
The issues we planned to address were to display a working calendar and possibly the ability to add and remove events.
We were able to succesfully display a working calendar, but adding and removing events can only be done through the code.
We are still learning the basics of Angular, Typescript, and the Full Calendar library. We are also learning to make it look
better on the actual website.

Back End:
The issues we planned to address were to form a connection with a database and determine what data we needed and how we wanted to structure that data within the database. Additionally, we were hoping to facilitate an interaction from client -> server -> database and back as a preliminary step to future scaling of the project. We were able to successfully decide on a database (mongoDB) and integrate it into our project. Additionally, we were able to successfully impliment a user registration page that accepts a username and password, checks both against criteria, hashes the password, and uploads them to the databse. Additionally, we were able to reverse the connection by checking to see that a new username was unique. This sprint we were not able to complete a similar process for creating events. This is because once we started development we realized that it made more sense to establish a user registration system before creating events for the sake of organization. However, we were able to plan out a general schema for the events and connect that to the database. 
