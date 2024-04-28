/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var prev *ListNode
    curr,flag := head,false 
    for curr != nil && curr.Next != nil {
        for curr != nil && curr.Next != nil && curr.Val == curr.Next.Val {
            curr = curr.Next
            flag = true
        }
        if flag {
            if prev == nil {
                head = curr.Next
                curr = head
            }else{
                prev.Next = curr.Next
                curr = prev
            }
            flag = false
        }else{
            prev = curr 
            curr = curr.Next
        }
    }
    return head
}