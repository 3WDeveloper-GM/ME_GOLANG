package main

import (
  "fmt"
  fv "finalv/method"
)

func main() {
  a := fv.CreateSlice(8)
  fmt.Println(a)
  fmt.Println(fv.FinalValueAfterOperations(a))
}
