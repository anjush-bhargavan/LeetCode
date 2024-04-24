/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    curr := head
    n := 1
    for curr != nil && n < k {
        curr = curr.Next
        n++
    } 
    temp := curr
    curr1 := head
    for curr.Next != nil {
        curr = curr.Next
        curr1 = curr1.Next
    }
    temp.Val,curr1.Val = curr1.Val, temp.Val
    return head

}