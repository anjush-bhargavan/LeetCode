/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    inOrder(root,&sum)
    return root
}

func inOrder(node *TreeNode,sum *int) {
    if node == nil {
        return
    }
    inOrder(node.Right,sum)
    node.Val += *sum
    *sum = node.Val
    inOrder(node.Left,sum)
}