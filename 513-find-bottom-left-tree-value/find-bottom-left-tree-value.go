/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    maxCount,count,min := -1,0,0
  
    LeftVal(root,&count,&maxCount,&min)

    return min
}


func LeftVal(node *TreeNode,count,maxCount,min *int) {
    if node.Left == nil && node.Right == nil {
        if *count > *maxCount {
            *min = node.Val
            *maxCount = *count
        }
        return
    }

    *count++
    if node.Left != nil && node.Right != nil  {
        LeftVal(node.Left,count,maxCount,min)
        LeftVal(node.Right,count,maxCount,min)
    }else if node.Left != nil {
         LeftVal(node.Left,count,maxCount,min)
    }else{
        LeftVal(node.Right,count,maxCount,min)
    }
   *count--
    
}