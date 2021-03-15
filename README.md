# Jikko-Golang

This small project is an HTTP web server on port 8080 that exposes endpoints for the purpose of exercise.

## Previous requirements

A MySQL database server should be running locally with a **jikkodb** database, which should have the **user** table and a **jikkouser** user identified by **jikkopass**.
It is possible to import the **jikkodb.sql** script on the **jikkodb** database to create and fill the **user** table.

## Compilation and Execution

Run `go build` to compile, and then `./jikko-golang` to run the server

## Endpoints

### /arrays/

This endpoint serves POST requests with the following JSON body:

```json
{
	"unsorted": [9, 8, 7, 6, 5, 4, 3, 2, 1, 0]
}
```

This should give a 200-status JSON response of

```json
{
	"unsorted": [9, 8, 7, 6, 5, 4, 3, 2, 1, 0],
	"sorted": [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
}
```

If there is any missing or incorrect data, it will return a 400 status code with the message:

> **Incorrect or missing data in request body**

### /users/

This endpoint serves either GET/POST requests, responding with an HTML page that displays a form to search user by id, and lists all the users in the database below

### /users/{id}

This endpoint serves either GET/POST requests, responding with an HTML page that displays the user detail page for the user identified by the provided index **id** (integer) as a URL parameter
