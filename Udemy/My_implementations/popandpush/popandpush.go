package popping

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Length int
	Head   *Node
}

func LinkedListInitialize() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) InsertatTail(data int) {

	newData := &Node{Data: data}
	if ll.Head == nil {
		ll.Head = newData
	} else {
		temp := ll.Head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = newData
	}
	ll.Length += 1
}

func (ll *LinkedList) InsertatHead(data int) {
	temp := ll.Head
	newData := &Node{Data: data}
	if ll.Head == nil {
		ll.Head = newData
	} else {
		newData.Next = temp
		ll.Head = newData
	}
	ll.Length += 1
}

func (ll *LinkedList) InsertatNode(data int, position int) {
	newData := &Node{Data: data}
	if ll.Head == nil {
		ll.InsertatHead(newData.Data)
	} else if position-1 == ll.Length {
		ll.InsertatTail(newData.Data)
	} else {
		temp := ll.Head
		temp2 := temp.Next
		for i := 0; i <= position-2; i++ {
			temp = temp.Next
			temp2 = temp2.Next
		}
		temp.Next = newData
		newData.Next = temp2
	}
	ll.Length += 1
}

func (ll *LinkedList) DeleteatHead() {
	if ll.Head != nil {
		temp := ll.Head
		ll.Head = temp.Next
		temp = nil
	}
	ll.Length -= 1
}

func (ll *LinkedList) DeleteatTail() {
	if ll.Head != nil {
		temp := ll.Head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp = nil
	}
	ll.Length -= 1
}

func (ll *LinkedList) DeleteatNode(position int) {
	if ll.Head != nil {
		temp := ll.Head
		temp2 := temp.Next
		for i := 0; i < position-1; i++ {
			temp = temp.Next
			temp2 = temp2.Next
		}
		temp.Next = temp2.Next
		temp2 = nil
	}
	ll.Length -= 1
}

func (ll *LinkedList) PrintL() {
	temp := ll.Head
	for temp.Next != nil {
		fmt.Print(temp.Data)
		temp = temp.Next
	}
	fmt.Println(temp.Data)
}
