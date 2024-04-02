/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    arr := []int{}
    sum,total := 0,0
    PathSums(root,&arr,&sum)
    for _,s := range arr {
        total += s
    }
    return total
}


func PathSums(root *TreeNode,arr *[]int,sum *int) {
    if root != nil {
        if root.Left == nil && root.Right == nil {
            *sum = (*sum * 10) + root.Val
            *arr = append(*arr,*sum)
            *sum /= 10
        }
        *sum = (*sum * 10) + root.Val
        PathSums(root.Left,arr,sum)
        PathSums(root.Right,arr,sum)
        *sum /= 10   
    }
}