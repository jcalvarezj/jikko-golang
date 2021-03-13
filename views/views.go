package views

import (
	"fmt"
	"net/http"

	"jikko-golang/sorting"
)

func HomeView(writer http.ResponseWriter, request *http.Request) {
	cosa := []int {9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Fprintf(writer, "<b>Welcome to the Home Page! 🐲</b>\n")
	fmt.Fprintf(writer, "<p>Sin ordernar: %v</p>", cosa)
	sorting.Mergesort(cosa)
	fmt.Fprintf(writer, "<p>Ordenado: %v</p>", cosa)
}
