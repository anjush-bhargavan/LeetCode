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
        return &TreeNode{Val:val,Left:root}
    }
    if d > depth  {
        return root
    }
    if d  == depth {
        root.Left = &TreeNode{Val:val,Left:root.Left}
        root.Right = &TreeNode{Val:val,Right:root.Right}
        return root
    }
    d++
    root.Left = addOneRow(root.Left, val, depth)
    root.Right = addOneRow(root.Right, val, depth)
    d--
    return root
}