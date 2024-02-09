/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    curr := head
    for curr != nil && curr.Next != nil {
        prev := curr
        curr = curr.Next
        newNode := &ListNode{Val : GCD(curr.Val,prev.Val)}
        prev.Next = newNode
        newNode.Next = curr   
    }
    return head
}

func GCD(a,b int) int {
    min := 0
    if a > b {
        min = b
    }else{
        min = a
    }
    for i := min ; i > 1 ; i-- {
        if a%i == 0 && b%i == 0 {
            return i
        }
    }
    return 1
}