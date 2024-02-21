/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    curr := head
    sum := 0
    temp := curr.Next
    for curr != nil && curr.Next != nil {
        if temp != nil && temp.Val != 0{
            sum += temp.Val
        }else{
            curr.Val = sum
            sum = 0
            if temp.Next != nil {
                curr.Next = temp
            }else{
                curr.Next = nil 
                break
            }
            curr = curr.Next
        }
        temp = temp.Next
    }
    return head
}