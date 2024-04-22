/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil {
        return head
    }
    curr,tail := head,head
    length := 0
    for curr != nil {
        tail = curr
        curr = curr.Next
        length++
    }

    k = k % length
    curr = head
    for _ = range length - k - 1 {
        curr = curr.Next
    }
    
    tail.Next = head
    head = curr.Next
    curr.Next = nil
    return head
}