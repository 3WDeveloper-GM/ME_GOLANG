package graph

import (
	"fmt"

	"github.com/gammazero/deque"
)

type Set struct {
	Elements map[string]bool
}

type AdjacencyList struct {
	Transformation map[string]Set
}

func InitializeSet() *Set {
	set := make(map[string]bool)
	return &Set{Elements: set}
}

func (s *Set) AddElement(Hash string) {
	s.Elements[Hash] = true
}

func InitializeAdjacencyList(adj map[string][]string) *AdjacencyList {
	ad := make(map[string]Set)
	for key, value := range adj {
		// generate set
		temp_set := InitializeSet()
		for _, vals := range value {
			temp_set.AddElement(vals)
		}
		ad[key] = *temp_set
	}
	return &AdjacencyList{Transformation: ad}
}

func (ad *AdjacencyList) Print() {
	for key, value := range ad.Transformation {
		fmt.Printf("%v : ", key)
		for key := range value.Elements {
			fmt.Printf("%v ", key)
		}
		fmt.Println("")
	}
}

func DFS(graph AdjacencyList, disc_map map[string]bool, root string) {
	fmt.Printf("Initizalizing DFS algorithm on vertex %v.\n", root)
	disc_map[root] = true
	fmt.Println(root)
	for key := range graph.Transformation[root].Elements {
		if _, ok := disc_map[key]; !ok {
			DFS(graph, disc_map, key)
		}
	}
}

func BFS(graph AdjacencyList, disc_map map[string]bool, root string) {
	var Q deque.Deque[string]
	fmt.Printf("Initializing BFS in the node %v.\n", root)
	fmt.Printf("%v", root)
	disc_map[root] = true
	Q.PushBack(root)
	for Q.Len() != 0 {
		v := Q.PopBack()
		for keys := range graph.Transformation[v].Elements {
			if _, ok := disc_map[keys]; !ok {
				fmt.Printf("%v", keys)
				disc_map[keys] = true
				Q.PushBack(keys)
			}
		}
	}
	fmt.Println("")
}

func EulerianPath(graph AdjacencyList, queue *deque.Deque[string], root string) {
	for keys := range graph.Transformation[root].Elements {
		if len(graph.Transformation[root].Elements) != 0 {
			delete(graph.Transformation[root].Elements, keys)
			EulerianPath(graph, queue, keys)
		} else {
			break
		}
	}
	queue.PushBack(root)
}
