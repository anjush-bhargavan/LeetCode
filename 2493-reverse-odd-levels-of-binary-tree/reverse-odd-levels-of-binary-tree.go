/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
    level := 0
    Reverser(root.Left,root.Right,&level)
    return root
}

func Reverser(node1,node2 *TreeNode,level *int) {
    if node1 == nil || node2 == nil {
        return
    }
    (*level)++
    if (*level) % 2 == 1 {
        node1.Val,node2.Val = node2.Val,node1.Val 
    }
    Reverser(node1.Left,node2.Right,level)
    Reverser(node2.Left,node1.Right,level)
    (*level)--
}