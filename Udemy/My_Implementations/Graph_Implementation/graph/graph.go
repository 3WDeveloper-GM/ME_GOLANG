package graph

import (
	"fmt"
	"implementation/set"
)

type Graph struct {
	Mapping map[string]set.Set
}

func Initialize() *Graph {
	mapping := make(map[string]set.Set)
	return &Graph{Mapping: mapping}
}

func (g *Graph) createNode(Hash string, s *set.Set) {
	g.Mapping[Hash] = *s
}

func (g *Graph) CreateGraph(init map[string][]string) {
	for key, values := range init {
		temp_set := set.Initialize()
		temp_set.CreateSet(values)
		g.createNode(key, temp_set)
	}
}

func (g *Graph) Print() {
	fmt.Println("{")
	for keys := range g.Mapping {
		fmt.Print("  ")
		fmt.Printf("%v: ", keys)
		for key := range g.Mapping[keys].Elements {
			fmt.Printf("%v ", key)
		}
		fmt.Println()
	}
	fmt.Println("}")
}
