/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    nodeMap := make(map[*ListNode]bool)
    curr := head
    for curr != nil && !nodeMap[curr] {
        nodeMap[curr] = true
        curr = curr.Next
    }
    return curr
}