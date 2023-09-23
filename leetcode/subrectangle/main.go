package main

import (
	"fmt"
	sr "subs/classsolution"
)

func main() {
  a := [][]int{{1,2,1},{4,3,4},{3,2,1},{1,1,1}}
  fmt.Println(a)
  b := sr.Constructor(a)
  b.Print()
  b.GetValue(0,2)
  b.UpdateSubrectangle(0,0,3,2,5)
  b.GetValue(0,2)
  b.GetValue(3,1)
  b.UpdateSubrectangle(3,0,3,2,10)
  b.GetValue(3,1)
  b.GetValue(0,2)
}
