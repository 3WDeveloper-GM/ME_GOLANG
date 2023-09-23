package main 

import (
  "fmt"
  ar "arconcat/method"
)

func main() {
  a := ar.CreateArray(12,13)
  b := ar.GetConcatenation(a)
  fmt.Println(a)
  fmt.Println(b)
}
