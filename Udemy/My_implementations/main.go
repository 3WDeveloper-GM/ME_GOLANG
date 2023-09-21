package main

import (
	"fmt"
	"math/rand"
	pop "pop/popandpush"
)

func main() {
	list := pop.LinkedListInitialize()
	for i := 0; i < 5; i++ {
		list.InsertatTail(rand.Intn(6))
	}
	list.PrintL()
	fmt.Println(list.Length)
	list.InsertatNode(2, 2)
	list.PrintL()
	list.DeleteatHead()
	list.PrintL()
	fmt.Println(list.Length)
	list.DeleteatNode(3)
	list.PrintL()
	fmt.Println(list.Length)
}
