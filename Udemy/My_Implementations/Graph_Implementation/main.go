package main

import (
	"implementation/graph"
	"implementation/set"
)

// Just testing the implementation of a basic graph structure in Go.
// I based the implementation using a basic "set" structure, that
// has fast lookup times and takes no memory other than that that is
// used to store the keys in a map.

func main() {
	a := map[string][]string{
		"a": {"b", "c"},
		"b": {"c", "d"},
		"c": {},
		"d": {"e", "f"},
		"e": {},
		"f": {"d"},
	}
	s1 := set.Initialize()
	s1.CreateSet(a["a"])
	s1.Print()
	g1 := graph.Initialize()
	g1.CreateGraph(a)
	g1.Print()
}
