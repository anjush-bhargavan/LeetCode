/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    mid := Middle(head)
    second := mid.Next
    mid.Next = nil

    secondReversed := Reverse(second)
    
    for head != nil && secondReversed != nil {
        if head.Val != secondReversed.Val {
            return false
        }
        head = head.Next
        secondReversed = secondReversed.Next
    }
    return true
}

//Middle function is to find  the middle node of the linked list.
func Middle (head *ListNode) *ListNode {
    slow,fast := head,head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

//Reverse function will reverse the given linked list and returns it.
func Reverse(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    curr := head.Next
    head.Next = nil
    for curr != nil {
        curr.Next,curr,head = head,curr.Next,curr
    }
    return head
}