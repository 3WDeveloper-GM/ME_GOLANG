package set

import "fmt"

// I recall that an empty struct takes no memory in Go.
// Perfect to implement a set structure in the mathematical
// sense of the word.
type Void struct{}

// A set is a map with keys of type string (this ones need memory)
// and values of type Void (which is an empty struct), that takes
// almost no memory
type Set struct {
	Elements map[string]Void
}

// Just a creator function, this one is just a function that
// initializes the "Set" struct.
func Initialize() *Set {
	elements := make(map[string]Void)
	return &Set{Elements: elements}
}

// Some methods for adding and removing members in the Set,
// AddMember adds members to the Set class and RemoveMember
// removes the key from the set.
func (s *Set) AddMember(Hash string) {
	var member Void
	s.Elements[Hash] = member
}

func (s *Set) RemoveMember(Hash string) {
	delete(s.Elements, Hash)
}

// This is a generic method to make a set using a large amount
// of elements in a slice of strings to create a big set. This one
// is useful, in particular, to create sets in a programatic
// manner. Is used for creating sets in the graph implementation
func (s *Set) CreateSet(elements []string) {
	for _, value := range elements {
		s.AddMember(value)
	}
}

// The following functions implement the two basic operations that
// can be done with sets, namely Union and Intersection.
func Union(Set1, Set2 *Set) *Set {
	new_set := Initialize()
	new_set.Elements = Set1.Elements
	for keys := range Set2.Elements {
		if _, ok := Set1.Elements[keys]; !ok {
			new_set.AddMember(keys)
		}
	}
	return new_set
}

func Intersection(s1, s2 *Set) *Set {
	new_set := Initialize()
	for keys := range s1.Elements {
		new_set.AddMember(keys)
	}
	for keys := range new_set.Elements {
		if _, ok := s2.Elements[keys]; !ok {
			new_set.RemoveMember(keys)
		}
	}
	return new_set
}

// Just a print function. It does exactly what it says.

func (s *Set) Print() {
	fmt.Print("{ ")
	for keys := range s.Elements {
		fmt.Printf("%v ", keys)
	}
	fmt.Println("}")
}
