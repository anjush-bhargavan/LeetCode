/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    var arr  []int
    result := [][]int{}
    sum := 0
    calcSum(root,&sum,&targetSum,&arr,&result)
    return result
}

func calcSum(root *TreeNode,sum,target *int,arr *[]int,result *[][]int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        *sum += root.Val
        *arr = append(*arr,root.Val)
        if *sum == *target {
            ar := make([]int, len(*arr))
            copy(ar, *arr)
            *result = append(*result,ar)
        }
        *arr = (*arr)[:len(*arr)-1]
        *sum -= root.Val
        return
    }
    *arr = append(*arr,root.Val)
    *sum += root.Val
    calcSum(root.Left,sum,target,arr,result)
    calcSum(root.Right,sum,target,arr,result)
    *sum -= root.Val
    *arr = (*arr)[:len(*arr)-1]
}