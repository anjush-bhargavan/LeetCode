/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    count,sum := 1,0
    d := depth(root)

    LeavesSum(root,&count,&sum,d)
    return sum
}

func LeavesSum(node *TreeNode,count,sum *int,d int) {
    if node == nil {
        return
    }
    if *count == d {
        fmt.Println(*count,node.Val)
        *sum += node.Val
    }
    *count++
    LeavesSum(node.Left,count,sum,d)
    LeavesSum(node.Right,count,sum,d)
    *count--        
}

func depth(node *TreeNode) int{
    if node == nil {
        return 0
    }
    left := depth(node.Left)
    right := depth(node.Right)

    if left < right {
        return right + 1
    }
    return left + 1
}