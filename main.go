// Entry point

package main

import (
	"fmt"
	"log"
	"net/http"

	"jikko-golang/arrays"

	"github.com/gorilla/mux"
)

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	cosa := []int {9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Fprintf(writer, "<b>Welcome to the Home Page! 🐲</b>\n")
	fmt.Fprintf(writer, "<p>Unordered: %v</p>\n", cosa)
	arrays.Mergesort(cosa)
	fmt.Fprintf(writer, "<p>Ordered: %v</p>\n", cosa)
}

func route() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/arrays", arrays.ArraysHandler).Methods("POST")

	fmt.Println("Started server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	route()
}
