func deckRevealedIncreasing(deck []int) []int {
    result := make([]int,len(deck))
    sort.Ints(deck)
    q := &Queue{}
    for i := range len(deck) {
        q.Push(i)
    } 
    for _,v := range deck {
        i := q.Pop()
        result[i] = v
        if q.Front != nil {
            i = q.Pop()
            q.Push(i)
        }
    
    }   
    return result
}


type Node struct {
    Val  int
    Next *Node
}
type Queue struct {
    Front *Node
    Rear  *Node
}

func (q *Queue) Push(data int) {
    newNode := &Node{Val:data}
    if q.Front == nil {
        q.Front = newNode
        q.Rear = newNode
        return
    }
    q.Rear.Next = newNode
    q.Rear = newNode
}

func (q *Queue) Pop() int {
    out := q.Front.Val
    q.Front = q.Front.Next
    return out
}