func evalRPN(tokens []string) int {
    s := &Stack{}
    for i := 0 ; i < len(tokens) ; i++ {
        if isOperator(tokens[i]) {
            b := s.Pop()
            a := s.Pop()
            switch tokens[i] {
            case "+":
                c := a + b
                s.Push(c)
            case "-":
                c := a - b
                s.Push(c)
            case "*":
                c := a * b
                s.Push(c)
            case "/":
                c := a / b
                s.Push(c)
            default:
                break
            }
        }else{
            data,_ := strconv.Atoi(tokens[i])
            s.Push(data)
        }
    }
    return  s.Pop()   
}

type Node struct {
    Val int
    Next *Node
}

type Stack struct {
    Top *Node
}

func(s *Stack) Push(data int) {
    newNode := &Node{ Val : data }
    newNode.Next = s.Top
    s.Top = newNode
}

func(s *Stack) Pop() int {
    data := s.Top.Val
    s.Top = s.Top.Next
    return data
}

func isOperator(s string) bool {
    if s == "+" || s == "-" || s == "*" || s == "/" {
        return true
    }
    return false
}