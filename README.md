# ChatAPP

## To View What's Currently There 

- to run the server: `go run src/.`
- to create the server's image: `docker compose up`

Either way, You can then access the application on **localhost**

## .ENV 

For this backend to work, you need to create your own .env file, in src/
it needs these variables to funciton: 
```shell
SERVER_ADDRESS="127.0.0.1:80"
GIN_MODE="DEBUG" //Optional
DATETIME_LAYOUT="2006-01-02T15:04:05.000Z"
URI_MONGODB="${mongoDbConnectionString}"
DB_NAME="ChatApp"
AUTH_DB_NAME="ChatAppAuth"
MAX_DB_CONNEXION_AGE="3600"
```

## Project's Objectives

- Be able to connect with an account to the app
- Be Able to communicate with any other user of the app, with a history
- Be Able to have multiple concurrent conversations
- Recieve messages on real time, through either Websockets or SSEs
- Deploy this app to the cloud, or at least to GitHub Pages
- Be able to use this app with HTTPS

## Current TODOs (By Priority)
- Store and fetch the conversations in mongoDb
    - add a CRUD for conversations
    - Add an endpoint to get the conversations of a user
    - From a conversation, add a way to get messages with pagination
    - create a way to add a message to a conversation
- Add An auth middleware using mongoDB
- Finalise the DI on the DB, so that we can plug it out if needed
- link the front to the backend
    - use queryparams to get messages from a conv, conv from a user
    - signup and login must be functional
- Beatufiy the CSS when a user send multiple messages
- Add a .env for the environment variables (currently in constants at top of their respective files)
- Add some unit tests