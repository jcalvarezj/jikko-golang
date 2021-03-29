// This file implements sorting algorithms

package arrays

// merge combines the left and right solutions in-place on the original array
func merge(array []int, left []int, right []int) {
	i, j, k := 0, 0, 0
	N, M := len(left), len(right)

	for i < N && j < M {
		if left[i] < right[j] {
			array[k] = left[i]
			i++
		} else {
			array[k] = right[j]
			j++
		}
		k++
	}

	for i < N {
		array[k] = left[i]
		i++
		k++
	}

	for j < M {
		array[k] = right[j]
		j++
		k++
	}
}

// Mergesort sorts in-place the input array using the Divide and Conquer method
func Mergesort(array []int) {
	N := len(array)
	if N > 1 {
		mid := N / 2

		left := array[:mid]
		right := array[mid:N]

		leftC := make([]int, len(left))
		rightC := make([]int, len(right))

		copy(leftC, left)
		copy(rightC, right)

		Mergesort(leftC)
		Mergesort(rightC)

		merge(array, leftC, rightC)
	}
}
