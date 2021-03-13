package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeView(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<b>Bienvenido al portal de inicio!!</b>")
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeView)
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("It works")

}
