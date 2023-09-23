package final

import "math/rand"

func CreateSlice(size int) []string {
  arr := make([]string, size)
  for index := range arr {
    switch {
    case rand.Intn(3) == 0:
      arr[index] = "X++"
    case rand.Intn(3) == 1:
      arr[index] = "X--"
    case rand.Intn(3) == 2:
      arr[index] = "--X"
    default:
      arr[index] = "++X"
    }
  }
  return arr
}

func FinalValueAfterOperations(operations []string) int {
  counter, lenn := 0 , len(operations)
  for i := 0; i < lenn; i++ {
    if operations[i] == "X++" || operations[i] == "++X" {
      counter ++
    } else if operations[i] == "X--" || operations[i] == "--X" {
      counter -- 
    }
  }
  return counter
}
