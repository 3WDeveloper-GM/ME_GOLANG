package main

import "fmt"

// I generate a map for fast lookup times
func arraytoDict(array []int) map[int][]int {
	dict := make(map[int][]int)
	for index, value := range array {
		dict[value] = append(dict[value], index)
	}
	return dict
}

// Create the groups of people according to the
// capacity.
func dictToarray(dict map[int][]int) [][]int {
	new_array := make([][]int, 0, len(dict))
	counter := 0
	for key, value := range dict {
		if key == len(value) {
			new_array = append(new_array, [][]int{value}...)
			counter++
		} else {
			for i := 0; i < len(value); i += key {
				end := i + key

				// necessary check to avoid slicing beyond
				// slice capacity
				if end > len(value) {
					end = len(value)
				}

				new_array = append(new_array, [][]int{value[i:end]}...)
			}

		}
	}
	return new_array
}

//func Groouppeople(groupsizes []int) [][]int {

//}

func main() {
	a := []int{1, 3, 3, 3, 2, 3, 2, 3, 3, 2, 2, 6, 6, 6, 6, 6, 6, 4, 4, 4, 4, 4, 4, 4, 4}
	fmt.Println(arraytoDict(a))
	fmt.Println(dictToarray(arraytoDict(a)))
}
