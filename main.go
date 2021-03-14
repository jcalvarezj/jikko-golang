// Entry point

package main

import (
	"fmt"
	"log"
	"net/http"

	"jikko-golang/arrays"
	"jikko-golang/users"

	"github.com/gorilla/mux"
)

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	array := []int {9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Fprintf(writer, "<b>Welcome to the Home Page! 🐲</b>\n")
	fmt.Fprintf(writer, "<p>Unordered: %v</p>\n", array)
	arrays.Mergesort(array)
	fmt.Fprintf(writer, "<p>Ordered: %v</p>\n", array)
}

func route() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/arrays", arrays.ArraysHandler).Methods("POST")
	router.HandleFunc("/users", users.UsersRootHandler)
	router.HandleFunc("/users/{id}", users.UserDetailHandler)

	fmt.Println("Started server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	route()
}
