/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    prev,slow,fast := head,head,head.Next

    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    if fast != nil {
        slow.Next = slow.Next.Next
    }else{
        prev.Next = slow.Next
    }

    return head
}
