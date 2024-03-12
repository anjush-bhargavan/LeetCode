/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    sumMap := make(map[int]*ListNode)
    temp := &ListNode{}
    temp.Next = head
    curr := temp
    sum := 0
    for curr != nil {
        sum += curr.Val
        sumMap[sum]= curr
        curr = curr.Next
    }
    sum = 0
    curr = temp
    for curr != nil {
        sum += curr.Val
        curr.Next = sumMap[sum].Next
        curr = curr.Next
    }
    return temp.Next
}
