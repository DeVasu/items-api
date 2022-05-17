# items-api
An API for CRUD Operations on any Product

## Written in Golang with Gorrila Mux and SQL Database

## Pre Requesites
* You should be connected to a live SQL Server (local server or remote server)
* settings for SQL connection can be done in `./datasources/items_db`
* Table `items` should already be existing in the database (setup in step 2)

## Setting up
* Run `go get .` in the application root folder
* Run `go run main.go`