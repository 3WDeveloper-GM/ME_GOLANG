package addtwo

type ListNode struct {
  Data int
  Next *ListNode
}

func (ln *ListNode) getLen() int {
  length, temp := 0, ln
  for temp != nil {
    temp = temp.Next
    length ++
  }
  return length
}

func Addtwonumbers(ln1, ln2 *ListNode) *ListNode {
  carry, sum := 0,0
  ln3 := new(ListNode)
  for ln1 != nil && ln2 != nil {
    sum = ln1.Data + ln2.Data + carry
    if sum >= 10{
      carry = 1
    }else {
      carry = 0
    }
    ln3.Data = sum % 10
    ln1 = ln1.Next
    ln2 = ln2.Next
  }
  return ln3.Next
}
