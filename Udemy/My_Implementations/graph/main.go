package main

import (
	"fmt"

	graph "itinerary/method"

	"github.com/gammazero/deque"
)

func main() {
	a := map[string][]string{
		"root": {"d"},
		"d":    {"a", "c"},
		"a":    {"root", "b"},
		"b":    {"c"},
		"c":    {"a", "d"},
	}
	g := graph.InitializeAdjacencyList(a)
	g.Print()
	disc_map := make(map[string]bool)
	graph.DFS(*g, disc_map, "a")
	disc_map = make(map[string]bool)
	graph.BFS(*g, disc_map, "a")
	fmt.Println("INITIALIZING EULERIAN PATH ALGORITHM")
	fmt.Println("///")
	queue := deque.New[string]()
	graph.EulerianPath(*g, queue, "root")
	for queue.Len() != 0 {
		fmt.Printf("%v ", queue.PopBack())
	}
	fmt.Println("")
	fmt.Println("///")
	fmt.Println("THE EULERIAN PATH IS PRINTED ABOVE.")
}
