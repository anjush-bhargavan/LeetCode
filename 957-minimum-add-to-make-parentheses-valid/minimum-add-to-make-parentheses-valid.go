func minAddToMakeValid(s string) int {
    st := &Stack{}
    for i := 0 ; i < len(s) ; i++ {
        if s[i] == '(' {
           st.Push(s[i])
        }else{
            if st.Peek() == '(' {
                st.Pop()
            }else{
                st.Push(s[i])
            }
        }
    }

    return st.Len()
}

type Node struct {
    val byte
    Next *Node
}

type Stack struct {
    Top *Node
    length int
}

func (s *Stack) Push(p byte) {
    newNode := &Node{val: p}
    newNode.Next = s.Top
    s.Top = newNode
    s.length++ 
}

func (s *Stack) Pop() {
    s.Top = s.Top.Next
    s.length--
}

func (s *Stack) Peek() byte {
    if s.Top == nil {
        return 0
    }
    return s.Top.val
}

func (s *Stack) Len() int {
    return s.length
}