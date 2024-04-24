/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    var temp *ListNode
    curr := head

    for curr != nil && curr.Val < x{
        temp = curr
        curr = curr.Next
    }
    prev := curr
    for curr != nil && curr.Next != nil {
        prev = curr
        curr = prev.Next
        if curr.Val < x {
            prev.Next = curr.Next
            if temp == nil {
                curr.Next = head
                head = curr
                temp = head
            }else{
                curr.Next = temp.Next
                temp.Next = curr
                temp = curr
            }
        }
    }
    return head
}