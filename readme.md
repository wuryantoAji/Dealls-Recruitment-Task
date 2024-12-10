# Technical Test - Software Engineer

Aji Wuryanto

## Service Structure
- database directory stores database package for initializing database, closing database, save user to database and get user from database
- model directory stores model package for user struct, function to hash the password and function to compare hash
- service directory stores service package for calling the database package whenever there is a request
- main.go serves as the root handler for the API request

## How to run the service
Start service by using
> go run .

Run unit test by using
> go test ./...

## Deployment


## Lint

