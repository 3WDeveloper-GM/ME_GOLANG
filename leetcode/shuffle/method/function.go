package shuffle

import "math/rand"

func CreateArray(size, randsize int) []int {
  new_array := make([]int, 2*size)
  for i := 0; i < 2*size;i++ {
    new_array[i] = rand.Intn(randsize)
  }
  return new_array
}

func Shuffle(nums []int, n int) []int {
  array := make([]int, 2*n)
  for i := 0; i<n ; i++{
    array[2*i] = nums[i]
    array[2*i+1] = nums[n+i]
  }
  return array
}
