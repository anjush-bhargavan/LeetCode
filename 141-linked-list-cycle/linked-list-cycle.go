/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    valMap := make(map[*ListNode]*ListNode)
    curr := head
    if curr == nil || curr.Next == nil {
        return false
    }
    for curr != nil {
        _, ok := valMap[curr]
        if ok && valMap[curr] == curr.Next{
            return true
        }else {
            valMap[curr] = curr.Next
        }
        curr = curr.Next
    }
    return false
}