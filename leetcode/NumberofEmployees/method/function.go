package function

import "fmt"

func ToDict(hours []int) map[int]int {
  dict := make(map[int]int)
  for index, val := range hours {
    dict[index] = val
  }
//  fmt.Printf("The new map is the following: %v.\n",dict)

  return dict
}

func NumberOfEmployeesWhoMetTarget(hours []int, target int) int {
  result := 0 
  dict := ToDict(hours)
  for _, element := range dict {
    if element >=  target {
      result ++
    }
  }
  fmt.Printf("The hours array is: %v \n", hours)
  fmt.Printf("The number of employes that reached the target of %v hours is: %v \n",target,result)
  return result 
}
