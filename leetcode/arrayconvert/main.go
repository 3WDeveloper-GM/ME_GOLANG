package main

import (
	"fmt"

	arc "arrayconv/method"
)

func main() {
	a := []int{1, 3, 4, 1, 2, 3, 1, 3, 3, 3}
	dict, _ := arc.CountElementArray(a)
	fmt.Println(arc.DictToarray(dict))
}
