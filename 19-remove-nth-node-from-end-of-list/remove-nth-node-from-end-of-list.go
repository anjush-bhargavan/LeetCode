/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    count,count1 := 1,1
    curr := head
    if  head.Next == nil {
        return nil
    }
    for curr != nil {
        curr = curr.Next
        count++
    }
    curr = head
    for curr != nil && (count-n-1) != count1{
        curr = curr.Next
        count1++
    }
    if curr == nil {
        return head.Next
    }
    if curr.Next!= nil {
        curr.Next = curr.Next.Next
    }
    return head
}