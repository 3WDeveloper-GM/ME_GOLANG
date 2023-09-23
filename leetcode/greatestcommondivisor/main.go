package main

import (
	"fmt"
	"math/rand"
)

type Listnode struct{
  Val int
  Next *Listnode
}

func GCD(a,b int) int {
  for b != 0 {
    t := b
    b = a%b
    a = t
  }
  return a
}

func (ln *Listnode) Print() {
  temp := ln
  for temp != nil {
    fmt.Print(temp.Val)
    temp = temp.Next 
  }
  fmt.Println()
}

func insertNext(head *Listnode) {
  temp := head
  temp.Next = &Listnode{Val: GCD(temp.Val, temp.Next.Val), Next: head.Next}
}

func insertGreatestCommonDivisors(head *Listnode) *Listnode {
  temp := head // get the head of the original list
  for temp.Next != nil {
    insertNext(temp)
    temp = temp.Next
    temp = temp.Next
  }
  return head
}

func main() {
  d := &Listnode{Val: rand.Intn(9)}
  c := &Listnode{Val: rand.Intn(9), Next: d}
  b := &Listnode{Val: rand.Intn(9), Next: c}
  a := &Listnode{Val: rand.Intn(9), Next: b}
  
  a.Print()
  
  insertGreatestCommonDivisors(a)

  a.Print()
}


