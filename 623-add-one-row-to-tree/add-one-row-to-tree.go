/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var d = 2
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if root == nil {
        return root
    }
    if depth == 1 {
            newNode := &TreeNode{Val:val}
            newNode.Left = root
            return newNode
    }
    if d > depth  {
        return root
    }
    if d  == depth {

            newNode1 := &TreeNode{Val:val}
            newNode1.Left = root.Left
            root.Left = newNode1

            newNode2 := &TreeNode{Val:val}
            newNode2.Right = root.Right
            root.Right = newNode2

        return root
    }
    d++
    _ = addOneRow(root.Left, val, depth)
    _ = addOneRow(root.Right, val, depth)
    d--
    return root
}