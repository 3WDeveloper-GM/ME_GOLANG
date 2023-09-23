package main

import (
  "fmt"
  sh "shuffle/method"
)

func main() {
  a := sh.CreateArray(3,12)
  b := sh.Shuffle(a, 3)
  fmt.Println(a)
  fmt.Println(b)
}
