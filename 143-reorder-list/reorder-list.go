/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    mid := Middle(head)
    second := mid.Next
    mid.Next = nil
    second = Reverse(second)
 
    curr := head
    for curr != nil && second != nil {
        curr.Next,curr = second,curr.Next
        second.Next,second = curr,second.Next
    }
}

// Middle function is to find the middle of the liked list.
func Middle (head *ListNode) *ListNode {
    slow,fast := head,head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    } 
    return slow
}

// Reverse function is to reverse the given linked list.
func Reverse (head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    curr := head.Next
    head.Next = nil
    for curr != nil {
        curr.Next,curr,head = head,curr.Next,curr
    }
    return head
}