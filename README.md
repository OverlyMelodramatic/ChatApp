# ChatAPP

## To View What's Currently There 

to run the server: `go run src/.`
to create the server's image: `docker compose up`

Either way, You can then access the application on **localhost**

## Projects Objectives

- Be able to connect with an account to the app
- Be Able to communicate with any other user of the app, with a history
- Be Able to have multiple concurrent conversations
- Recieve messages on real time, through either Websockets or SSEs
- Deploy this app to the cloud, or at least to GitHub Pages
- Be able to use this app with HTTPS

## Current TODOs
- Add an JWT token based auth to the backend
- Add a mongoDB database
- Beatufiy the CSS when a user send multiple messages
- Add a .env for the environment variables (currently in constants at top of their respective files)
- Add some unit tests