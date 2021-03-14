// Entry point

package main

import (
	"fmt"
	"log"
	"net/http"

	"jikko-golang/arrays"
	"jikko-golang/users"
	"jikko-golang/repository"

	"github.com/gorilla/mux"
)

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

	fmt.Println("Started server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Entry point
func main() {
	dbConn, err := repository.ConnectToDatabase()
	if err != nil {
		fmt.Println("Could not connect to database. 'users' won't work properly")
	}
	usersHandler := users.UsersHandler{DbConn: dbConn}
	defer dbConn.Close()
	route(&usersHandler)
}
