/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    prev := head

    if prev == nil {
        return prev
    }

    curr := prev.Next
    prev.Next = nil

    for curr != nil {
        temp := curr.Next
        curr.Next = prev
        prev,curr = curr,temp
    }

    return prev
}