//This file defines the models involved in this module

package arrays

// ArraysJSON represents the JSON response and request unit
type ArraysJSON struct {
	Unsorted []int `json:"unsorted"`
	Sorted []int `json:"sorted"`
}
