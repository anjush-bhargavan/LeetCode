/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return IsSame(root.Left,root.Right)
}

func IsSame(node1,node2 *TreeNode) bool {
    if node1 == nil || node2 == nil {
        return node1 == node2
    }
    if node1.Val != node2.Val {
        return false
    }

    return IsSame(node1.Left,node2.Right) && IsSame(node1.Right,node2.Left)
}