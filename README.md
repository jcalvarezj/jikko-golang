# Jikko-Golang

This small project is an HTTP web server on port 8080 that exposes endpoints for the purpose of exercise.

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

## Compilation and Execution

Run `go build` to compile, and then `./jikko-golang` to run the server
