package main

import (
	"fmt"
	"log"
	"net/http"

	"jikko-golang/arrays"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", arrays.HomeHandler)
	fmt.Println("Started server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
