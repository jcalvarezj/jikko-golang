// This file defines Handler view functions to be mapped to Web resources

package arrays

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

// ArraysHandler godoc
// @Summary Serves the /arrays resource POST requests
// @Description Serves the /arrays resource POST requests by receiving a JSON object with an unsorted numbers array and responds with a JSON object with that array and its sorted version
// @Param unsorted body arrays.ArraysJSON true "The unsorted array of numbers"
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 405 {object} string
// @Router /arrays [post]
func ArraysHandler(writer http.ResponseWriter, request *http.Request) {
	var jsonData ArraysJSON
	contents, err := ioutil.ReadAll(request.Body)
	json.Unmarshal(contents, &jsonData)

	array := jsonData.Unsorted

	if len(array) > 0 {
		if err != nil {
			fmt.Println("An error has occurred when parsing JSON data: ", err)
		} else {
			original := make([]int, len(array))
			copy(original, array)

			Mergesort(array)
			jsonData.Sorted = array
			jsonData.Unsorted = original

			jsonResponse, err2 := json.Marshal(&jsonData)

			if err2 != nil {
				fmt.Println("An error has occurred when packing JSON response", err)
			} else {
				writer.Header().Set("Content-Type", "application/json; charset=utf-8")
				writer.Write(jsonResponse)
			}
		}
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, "Incorrect or missing data in request body")
	}
}
