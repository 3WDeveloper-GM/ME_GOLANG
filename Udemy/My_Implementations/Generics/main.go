// I was learning generics in golang, and wanted to implement
// two simple functions that can be used for addition and substraction
// of multiple data types.
// This is an academic exercise.

package main

import (
	"fmt"
)

func additon[T ~int | ~float64 | ~string](a, b T) T {
	return a + b
}

func substraction[T ~int | ~float64 | ~float32](a, b T) T {
	return a - b
}

func main() {
	fmt.Println(additon[int](3, 2))
	fmt.Println(additon[float64](2.3, 1.3))
	fmt.Println(additon[string]("Hi, ", "mom"))
	fmt.Println(additon[string]("Hello, ", "World"))
	fmt.Println(substraction[float32](2.0000, 1.3))
}
