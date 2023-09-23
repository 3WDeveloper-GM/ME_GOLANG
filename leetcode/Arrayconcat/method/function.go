package arconcat

import "math/rand"

func CreateArray(size, randsize int) []int {
  new_array := make([]int, size)
  for index := range new_array {
    new_array[index] = rand.Intn(randsize)
  }
  return new_array
}

func GetConcatenation(nums []int) []int {
  lenn := len(nums)
  concatenatedarray := make([]int, 2*lenn)
  for i := 0 ; i < 2*lenn; i++ {
    concatenatedarray[i] = nums[i%lenn]
  }
  return concatenatedarray
}
