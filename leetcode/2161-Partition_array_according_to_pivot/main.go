package main

import "fmt"

func SortList(array []int, target int) []int {
	new_array := make([]int, len(array))
	new_array_index := 0
	// This sorts the list first by sorting the elements that are
	// lower than the target. This creates the lower half of the array
	for _, value := range array {
		if value < target {
			new_array[new_array_index] = value
			new_array_index++
		}
	}
	// If there are elements that are the same in the target, these go
	// in the middle, the tracking of the array is done by updating a
	// variable of type integer.
	for _, value := range array {
		if value == target {
			new_array[new_array_index] = value
			new_array_index++
		}
	}
	// Same as the first run, but with the second half of the array,
	// so it has the higher elements.
	for _, value := range array {
		if value > target {
			new_array[new_array_index] = value
			new_array_index++
		}
	}
	return new_array
}

// Testing my function
func main() {
	a := []int{9, 12, 5, 10, 14, 3, 10}
	fmt.Println(SortList(a, 10))
}
