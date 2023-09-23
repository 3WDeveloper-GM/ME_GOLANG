package method

func PivotArray(nums []int, pivot int) []int {
  length := len(nums)
  new_array := make([]int, length)
  new_array[length/2] = pivot
  counter1 := 0
  counter2 := 0
  for _, value := range nums {
    if value <= pivot {
      new_array[counter1] = value
      counter1 ++
    } else {
      new_array[length/2+counter2] = value
      counter2 ++
    }
  }
  return new_array
}
