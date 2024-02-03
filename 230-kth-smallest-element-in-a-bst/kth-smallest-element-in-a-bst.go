/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    arr := []int{}
    inorder(root,&arr)
    for i := 0 ; i < len(arr) ; i++ {
        if i == k-1 {
            return arr[i]
        }
    }
    return root.Val
}

func inorder(root *TreeNode, arr *[]int) {
    if root == nil {
        return
    }
    inorder(root.Left,arr)
    *arr = append(*arr,root.Val)
    inorder(root.Right,arr)
}