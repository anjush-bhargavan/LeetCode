/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    curr := head
    nodeMap := make(map[*Node]*Node)
    for curr != nil {
        copyNode := &Node{Val : curr.Val}
        nodeMap[curr] = copyNode
        curr = curr.Next
    }

    curr = head
    for curr != nil {
        copyNode := nodeMap[curr]
        copyNode.Next = nodeMap[curr.Next]
        copyNode.Random = nodeMap[curr.Random]
        curr = curr.Next
    }


    return nodeMap[head]
}