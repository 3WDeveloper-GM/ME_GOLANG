// This is the dumb brute force solution. It's time complexity it's
// O(n^2). I think that this solution might be useful for a starting
//point.

package main

import "fmt"

func Twosum(target int, array []int) []int {
	for index_1 := range array {
		for index_2 := range array {
			result := array[index_1] + array[index_2]
			if result == target {
				a := []int{index_1, index_2}
				return a
			}
		}
	}
	return []int{-1, -1}
}

func main() {
	a := []int{1, 3, 4, 9, 6, 8, 2, 6, 7}
	target := 10
	fmt.Println(Twosum(target, a))
}
