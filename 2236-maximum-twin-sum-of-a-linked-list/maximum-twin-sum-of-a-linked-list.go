/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    max := 0
    sum :=  TwinSum(head)
    for sum != nil {
        if max < sum.Val {
            max = sum.Val
        }
        sum = sum.Next
    } 
    return max  
}

func TwinSum(head *ListNode) *ListNode {
    mid := FindMiddle(head)

    sec := mid.Next
    mid.Next = nil
    rev := Reverse(sec)
    fmt.Println(rev)
    newList := &ListNode{}
    curr := newList
    for rev != nil {
        curr.Next =&ListNode{Val : rev.Val + head.Val}
        rev = rev.Next
        head = head.Next
        curr = curr.Next  
    }
    return newList.Next
}


func FindMiddle( node *ListNode) *ListNode {
    slow,fast := node,node.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    } 
    return slow
}

func Reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    newHead := Reverse(head.Next)

    head.Next.Next = head
    head.Next = nil

    return newHead
}