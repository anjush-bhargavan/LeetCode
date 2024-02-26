func dailyTemperatures(temperatures []int) []int {
result := make([]int,len(temperatures))
s := &Stack{}
for i := 0 ; i <  len(temperatures) ; i++ {
   for s.length > 0 && temperatures[s.Peek()] < temperatures[i] {
       index := s.Peek()
       s.Pop()
       result[index] = i - index
   }
    s.Push(i)
}
return result
}


type Node struct {
    Val int
    Next *Node
}

type Stack struct {
    Top *Node
    length int
}

func (s *Stack) Push(data int) {
    newNode := &Node{Val:data}
    newNode.Next = s.Top
    s.Top = newNode
    s.length++
}

func (s *Stack) Pop() {
    s.Top = s.Top.Next
    s.length--
}

func (s *Stack) Peek() int {
    if s.Top == nil {
        return 0
    }
    return s.Top.Val
}