//This is a better solution using hash maps, because the lookup is O(1) in time complexity
//this algorithm runs faster than the brute force solution. 


package main

import (
	"fmt"
	"math/rand"
	"time"
)

func TwoSumHash(array []int, target int) []int {
	hash := make(map[int]int)

	//Initialize Hash Map
	for index, value := range array {
		complement := target - array[index]
		if _, ok := hash[complement]; ok && hash[complement] != index {
			return []int{index, hash[complement]}
		}
		hash[value] = index
	}
	return []int{-1, -1}
}

func main() {
	n := 1000
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(20)
	}

	fmt.Println(a)
	now := time.Now()
	defer func() {
		fmt.Println(time.Now().Sub(now))
	}()
	fmt.Println(TwoSumHash(a, 6))
	fmt.Println(a)
}
