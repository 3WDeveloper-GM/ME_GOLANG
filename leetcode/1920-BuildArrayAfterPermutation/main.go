package main

import (
	"fmt"

	pm "permarray/method"
)

func main() {
	a := pm.RandArray(10, 9)
	fmt.Println(a)
	b := pm.BuildArray(a)
	fmt.Println(b)
}
