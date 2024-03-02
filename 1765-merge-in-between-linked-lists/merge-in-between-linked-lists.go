/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    count := 0
    curr := list1
    var first *ListNode
    for curr != nil && count - 1 != b{
        if count + 1 == a {
            first = curr
        }
        count++
        curr = curr.Next
    }

    first.Next = list2
    curr1 := first
    for curr1.Next != nil {
        curr1 = curr1.Next
    }
    curr1.Next = curr
    
    return list1
}