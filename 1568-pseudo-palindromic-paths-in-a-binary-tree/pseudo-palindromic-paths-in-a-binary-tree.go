/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths (root *TreeNode) int {
    count := make(map[int]int)
    res := PathCount(root,&count)
    return res
}


func PathCount(root *TreeNode,count *map[int]int) int {
    if root.Left == nil && root.Right == nil {
        odd := 0
        (*count)[root.Val]++

        for _,v := range (*count) {
            if v % 2 == 1 {
                odd++
            }
        }
        (*count)[root.Val]--

        if odd == 1 || odd == 0 {
            return 1
        }
        return 0
    }
    left,right := 0,0
    (*count)[root.Val]++

    if root.Left !=  nil {
        left = PathCount(root.Left,count)
    }
    if root.Right != nil {
        right = PathCount(root.Right,count)
    }
    
    (*count)[root.Val]--

    return left + right
}