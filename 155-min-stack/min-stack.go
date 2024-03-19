type Node struct{
    data int
    next *Node
}
type MinStack struct {
    top *Node
    Min *Node
}


func Constructor() MinStack {
    stack:= MinStack{}
    return stack
}


func (this *MinStack) Push(val int)  {
    newNode := &Node{data:val}
    minNode := &Node{data:val}
    newNode.next = this.top
    this.top = newNode
    if this.Min == nil {
        minNode.next = this.Min
        this.Min = minNode
    }else{
        if this.Min.data >= val {
            minNode.next = this.Min
            this.Min = minNode
        }
    }
}


func (this *MinStack) Pop()  {
    if this.top.data == this.Min.data {
        this.Min = this.Min.next
    }
    this.top = this.top.next
    
}


func (this *MinStack) Top() int {
    return this.top.data
}


func (this *MinStack) GetMin() int {
    return this.Min.data
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */