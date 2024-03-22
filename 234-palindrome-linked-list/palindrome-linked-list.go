/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    mid := Middle(head) //finds the middle
    second := mid.Next // second will be the head reference of the linkedlist after the middle. 
    mid.Next = nil     // disconnecting the linked list after the middle

    second = Reverse(second) //reversing the second part which we disconnected
    
    for head != nil && second != nil {
        if head.Val != second.Val {   //checking each node values are same or not. if not we return false
            return false
        }
        head = head.Next
        second = second.Next
    }
    return true // after one of them become nil we return true.
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