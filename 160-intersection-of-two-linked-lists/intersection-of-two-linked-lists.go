/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    nodeMap := make(map[*ListNode]*ListNode)

    for headA != nil  {
        nodeMap[headA] = headA.Next
        headA = headA.Next
    }
    for headB != nil {
        if _,ok := nodeMap[headB]; !ok {
            nodeMap[headB] = headB.Next
        }else{
            return headB
        }
        headB = headB.Next
    }
    return nil
}