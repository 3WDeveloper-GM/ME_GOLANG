package main

import (
	"fmt"
	"math"
	"math/rand"
)

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func numberofbits(n int) int {
	result := 0
	if n == 0 {
		result += 0
	} else if n == 1 {
		result += 1
	} else {
		result += 1 + numberofbits(n-IntPow(2, int(math.Log2(float64(n)))))
	}
	//fmt.Println(result)
	return result
}

func countBits(n int) []int {
	result := make([]int, n+1)
	for index := range result {
		result[index] = numberofbits(index)
	}
	return result
}

func main() {
	a := rand.Intn(19)
	fmt.Println(a)
	fmt.Println(numberofbits(a))
	//for i := 0; i <= 9; i++ {
	//		fmt.Println(numberofbits(i))
	//}
	fmt.Println(countBits(a))
}
