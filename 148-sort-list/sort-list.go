/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }   

    mid := Middle(head)
    sec := mid.Next
    mid.Next = nil

    left := sortList(head)
    right := sortList(sec)

    return Merge(left,right)
}

func Middle(node *ListNode) *ListNode {
    slow,fast := node,node.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

func Merge(node1,node2 *ListNode) *ListNode {
    node := &ListNode{}
    curr := node
    for node1 != nil && node2 != nil {
        if node1.Val < node2.Val {
            curr.Next = node1
            node1 = node1.Next
        }else{
            curr.Next = node2
            node2 = node2.Next
        }
        curr = curr.Next
    }

    if node1 != nil {
        curr.Next = node1
    }
    if node2 != nil {
        curr.Next = node2
    }
    return node.Next
}