/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    curr := head
    if curr == nil || curr.Next == nil {
        return head
    }
    head = curr
    prev := curr
    curr = curr.Next
    head = curr
    next := curr.Next
    prev.Next = next
    curr.Next = prev
    curr = prev.Next
    if curr == nil {
        return head
    }
    next = curr.Next

    for curr != nil && next != nil{
        temp := next.Next
        prev.Next = next
        next.Next = curr  
        curr.Next = temp
        prev = curr
        curr = temp
        if curr == nil {
            break
        }
        next = curr.Next

    }
    return head
}