package method

import "math/rand"

func RandArray(size, randsize int) []int {
  list := make([]int, size)
  for index := range list {
    list[index] = rand.Intn(randsize)
  }
  return list
}

func BuildArray(nums []int) []int {
  new_array := make([]int, len(nums))
  for index := range nums {
    new_array[index] = nums[nums[index]]
  }
  return new_array
}

