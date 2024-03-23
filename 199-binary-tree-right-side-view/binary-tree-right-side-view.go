/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    arr := []int{}
    count,max := 0,0
    curr := root
    for curr != nil {
        arr = append(arr,curr.Val)
        curr = curr.Right
        count++
    }
    max = count
    count = 1
    RightTree(root,&arr,&count,&max)
    return arr
}


func RightTree(node *TreeNode,arr *[]int,count,max *int) {
    if node == nil {
        return
    }
    if *max < *count {
        *arr = append(*arr,node.Val)
        *max = *count
    }
    *count++
    RightTree(node.Right,arr,count,max)
    RightTree(node.Left,arr,count,max)
    *count--
}