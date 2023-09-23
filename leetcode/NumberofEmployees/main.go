package main

import fn "number/method"

func main() {
  a := []int{1,2,2,2,3,1,5,4,6,2}
  fn.ToDict(a)
  fn.NumberOfEmployeesWhoMetTarget(a,2)
}
