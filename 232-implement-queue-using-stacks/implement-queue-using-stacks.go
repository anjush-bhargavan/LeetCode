type MyQueue struct {
    s1 *Stack
    s2 *Stack
}


func Constructor() MyQueue {
    return MyQueue{s1: &Stack{}, s2: &Stack{}}
}


func (this *MyQueue) Push(x int)  {
    this.s1.Push(x)
    return 
}


func (this *MyQueue) Pop() int {
    if !this.s2.IsEmpty() {
        data := this.s2.Peak()
        this.s2.Pop()
        return data
    }
    fmt.Println(this.s1.Peak())
    for !this.s1.IsEmpty() {
        this.s2.Push(this.s1.Peak())
        this.s1.Pop()
    }
    data := this.s2.Peak()
    this.s2.Pop()
    return data
}


func (this *MyQueue) Peek() int {
    if !this.s2.IsEmpty() {
        return this.s2.Peak()
    }

    for !this.s1.IsEmpty() {
        this.s2.Push(this.s1.Peak())
        this.s1.Pop()
    }
    return this.s2.Peak()
}


func (this *MyQueue) Empty() bool {
    return this.s1.IsEmpty() && this.s2.IsEmpty()
}

type Node struct{
    Val     int
    Next    *Node
}

type Stack struct {
    Top *Node
}

func(s *Stack) Push(data int) {
    newNode := &Node{Val:data}
    newNode.Next = s.Top
    s.Top = newNode
}

func (s *Stack) Pop() {
    s.Top = s.Top.Next
}

func (s *Stack) IsEmpty() bool {
    return s.Top == nil
}

func (s *Stack) Peak() int {
    return s.Top.Val
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */