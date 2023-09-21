package main

import "fmt"

func binarysearchalgorithm(array []int, target, lowIndex, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	mid := (lowIndex + highIndex) / 2
	if array[mid] > target {
		return binarysearchalgorithm(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarysearchalgorithm(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func binarysearchalgorithmfor(array []int, target, LowIndex, HighIndex int) int {
	start, end := LowIndex, HighIndex
	mid := (start + end) / 2
	for start <= end {
		if array[mid] > target {
			end = mid
		} else if array[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
		mid = (start + end) / 2
	}
	return -1
}

func main() {
	alpha := []int{1, 3, 6, 9, 10, 11, 12, 16, 20, 22, 26, 30, 32, 33, 60}
	fmt.Println(binarysearchalgorithm(alpha, 10, 0, len(alpha)))
	fmt.Println(binarysearchalgorithmfor(alpha, 10, 0, len(alpha)))
}
