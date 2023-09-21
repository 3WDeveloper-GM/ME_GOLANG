package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func candy(array []int) []int {
	na := make([]int, len(array))

	//Array initialization
	for index := range na {
		na[index] = 1
	}

	//Left pass
	for index := 1; index < len(na); index++ {
		if array[index] > array[index-1] {
			na[index] = na[index-1] + 1
		}
	}

	//right pass
	for index := len(na) - 2; index >= 0; index-- {
		if array[index] > array[index+1] {
			if na[index] <= na[index+1] {
				na[index] = na[index+1] + 1
			}
		}
	}

	return na
}

func main() {
	a := []int{1, 0, 2}
	fmt.Println(candy(a))
}
