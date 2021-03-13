//This file defines the models involved in this module

package arrays

type ArraysJSON struct {
	Unsorted []int `json: "unsorted"`
	Sorted []int `json: "sorted"`
}
