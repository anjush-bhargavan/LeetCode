/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    arr := [][]int{}
    level := 0
    Traversal(root,&arr,&level)
    return arr
}

func Traversal(root *TreeNode,arr *[][]int,level *int) {
    if root == nil {
        return
    }
    if len(*arr) <= *level {
        *arr = append(*arr,[]int{})
    }
    (*arr)[*level] = append((*arr)[*level],root.Val)
    *level++
    Traversal(root.Left,arr,level)
    Traversal(root.Right,arr,level)
    *level--
}