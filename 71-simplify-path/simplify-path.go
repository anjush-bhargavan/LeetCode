func simplifyPath(path string) string {
    s := &Stack{}
    dir := ""
    for i := 1 ; i < len(path) ; i++ {
        if path[i] != '/' {
            dir += string(path[i])
        }else {
            if dir == ".." {
                s.Pop()
            }else if dir != "" && dir != "."{
                dir = "/" + dir
                s.Push(dir)
            }
            dir = ""

        } 
    }
    if dir != "" {
            if dir == ".." {
                s.Pop()
                dir = ""
            }else if dir != "" && dir != "."{
                dir = "/" + dir
                s.Push(dir)
                dir = ""
            }
    }
    dir = ""
    for s.Top != nil {
        dir = s.Pop() +  dir
    }
    if dir == "" {
        return "/"
    }
    return dir
}


type Node struct {
    Val string
    Next *Node
}
type Stack struct {
    Top *Node
}

func(s *Stack) Push(data string) {
    newNode := &Node{Val : data }
    newNode.Next = s.Top
    s.Top = newNode
}

func (s *Stack) Pop() string {
    if s.Top == nil {
        return ""
    }
    str  := s.Top.Val
    s.Top = s.Top.Next
    return str
}
