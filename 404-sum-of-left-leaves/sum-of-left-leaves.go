/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var flag = false
func sumOfLeftLeaves(root *TreeNode) int {
    sum := 0
    LeafSum(root,&sum)
    return sum
}

func LeafSum(root *TreeNode,sum *int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        if flag {
            flag = false
            *sum += root.Val
            return 
        }
        return
    }
    flag = true
    LeafSum(root.Left,sum)
    flag = false
    LeafSum(root.Right,sum)
}