/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    arr := [][]int{}
    level := 0
    LevelOrder(root,&level,&arr)
    for i := 1 ; i < len(arr) ; i += 2 {
        Reverse(arr[i])
    }
    return arr
}

func Reverse(arr []int) {
    for i := 0 ; i < (len(arr)/2) ; i++ {
        arr[i],arr[len(arr)-1-i] = arr[len(arr)-1-i] ,arr[i] 
    }
} 

func LevelOrder(root *TreeNode,level *int,arr *[][]int) {
    if root == nil {
        return 
    }
    if len(*arr)  <= *level  {
        (*arr) = append(*arr,[]int{})
    }
    (*arr)[*level] = append((*arr)[*level],root.Val)
    *level++
    LevelOrder(root.Left,level,arr)
    LevelOrder(root.Right,level,arr)
    *level--
}