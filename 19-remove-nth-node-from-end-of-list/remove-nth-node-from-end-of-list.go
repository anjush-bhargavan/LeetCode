/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    curr1,curr2,prev := head,head,head
    for n > 0 {
        curr1 = curr1.Next
        n--
    }
    
    if curr1 == nil {
        return head.Next
    }

    for curr1 != nil {
        curr1 = curr1.Next
        prev = curr2
        curr2 = curr2.Next
    }

    prev.Next = curr2.Next

    return head
}