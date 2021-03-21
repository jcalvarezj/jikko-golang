// Entry point

package main

import (
	"fmt"
	"log"
	"net/http"

	_ "jikko-golang/docs"
	"jikko-golang/arrays"
	"jikko-golang/users"
	"jikko-golang/repository"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
)

// @title Jikko-Golang REST API
// @version 1.0
// @description This is an example Web server that sorts arrays and displays user info
// @host localhost:8080
// @BasePath /

// Handler view for the root resource
func homeHandler(writer http.ResponseWriter, request *http.Request) {
	array := []int {9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Fprintf(writer, "<b>Welcome to the Home Page! 🐲</b>\n")
	fmt.Fprintf(writer, "<p>Unordered: %v</p>\n", array)
	arrays.Mergesort(array)
	fmt.Fprintf(writer, "<p>Ordered: %v</p>\n", array)
}

// Registers resource mappings to Mux router and starts the server
func route(usersHandler *users.UsersHandler) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/arrays", arrays.ArraysHandler).Methods("POST")
	router.HandleFunc("/users", usersHandler.UsersRootHandler)
	router.HandleFunc("/users/{id}", usersHandler.UserDetailHandler)

	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	//router.HandleFunc("/swagger", httpSwagger.WrapHandler)
	router.NotFoundHandler = http.HandlerFunc(usersHandler.NotFoundHandler)

	fmt.Println("Started server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Entry point
func main() {
	dbConn, err := repository.ConnectToDatabase()
	if err != nil {
		fmt.Println("Could not create connection object")
		panic(err)
	}

	err = dbConn.Ping()
	if err != nil {
		fmt.Println("Couldn't connect to the database. 'users' won't work properly")
		panic(err)
	}

	fmt.Println("Connection to database was successful!")
	usersHandler := users.UsersHandler{DbConn: dbConn}
	defer dbConn.Close()
	route(&usersHandler)
}
