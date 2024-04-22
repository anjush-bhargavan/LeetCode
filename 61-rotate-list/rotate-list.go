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
    curr := head
    length := 0
    for curr != nil {
        curr = curr.Next
        length++
    }
    k = k % length
    curr = head
    for _ = range k {
        curr = curr.Next
    }
    curr1 := head
    for curr.Next != nil {
        curr = curr.Next
        curr1 = curr1.Next
    }
    curr.Next = head
    head = curr1.Next
    curr1.Next = nil
    return head
}