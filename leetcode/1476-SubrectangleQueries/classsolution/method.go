package classsolution

import "fmt"

type SubrectangleQueries struct {
	Row      int
	Column   int
	Elements [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	a := &SubrectangleQueries{Row: 0, Column: 0, Elements: nil}
	a.Row = len(rectangle)
	a.Column = len(rectangle[0])
	a.Elements = rectangle
  return *a
}

func (rec *SubrectangleQueries) Print() {
	fmt.Printf("This is the subrectangle class; it has %v rows and %v columns.\n",rec.Row, rec.Column)
  fmt.Println("///")
  fmt.Println("///")
  fmt.Println("The rectangle is:")
  for index := range rec.Elements {
    fmt.Println(rec.Elements[index])
  }
  fmt.Println("")
}

func (rec *SubrectangleQueries) GetValue(row, column int) int {
  result := rec.Elements[row][column]
  fmt.Printf("The element of the %v row and %v column is: %v.\n", row, column, result)
  return result
}

func (rec *SubrectangleQueries) UpdateRow(row, StartIndex, EndIndex, value int) {
  fmt.Println("We have changed a row of the rectangle object:")
  fmt.Println("///")
  fmt.Println("///")
  fmt.Println("The new rectangle is: ")
  for dummy := StartIndex ; dummy <= EndIndex; dummy ++ {
    //fmt.Println(dummy)
    rec.Elements[row][dummy] = value
  }
  for index := range rec.Elements {
    fmt.Println(rec.Elements[index])
  }
}


func (rec *SubrectangleQueries) updateRow(row, StartIndex, EndIndex, value int) {
  //fmt.Println("We have changed a row of the rectangle object:")
  //fmt.Println("///")
  //fmt.Println("///")
  //fmt.Println("The new rectangle is: ")
  for dummy := StartIndex ; dummy <= EndIndex; dummy ++ {
    //fmt.Println(dummy)
    rec.Elements[row][dummy] = value
  }
//  for index := range rec.Elements {
//    fmt.Println(rec.Elements[index])
//  }
}

func (rec *SubrectangleQueries) UpdateSubrectangle(StartRow, StartColumn, EndRow, EndColumn, value int) {
  fmt.Println("We have changed a subrectangle, the new subrectangle is:")
  fmt.Println("///")
  fmt.Println("///")
  for dummy := StartRow; dummy <= EndRow; dummy ++ {
    rec.updateRow(dummy,StartColumn, EndColumn, value)
  }
  for index := range rec.Elements {
    fmt.Println(rec.Elements[index])
  }

}
