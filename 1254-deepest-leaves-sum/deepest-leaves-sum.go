/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    count,sum,d := 1,0,0
    // d := depth(root)

    LeavesSum(root,&count,&sum,&d)
    return sum
}

func LeavesSum(node *TreeNode,count,sum,d *int) {
    if node == nil {
        return
    }
    if *count == *d {
        *sum += node.Val
    }
    if *count > *d {
        *d = *count
        *sum = node.Val
    }
    *count++
    LeavesSum(node.Left,count,sum,d)
    LeavesSum(node.Right,count,sum,d)
    *count--        
}
