/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == target && root.Left == nil && root.Right == nil {
        return nil
    }

  


    root.Left = removeLeafNodes(root.Left,target)
    root.Right = removeLeafNodes(root.Right,target)
    
    if root.Val == target && root.Left == nil && root.Right == nil {
        return nil
    }

    return root
}


// func IsLeafNode(root *TreeNode, target int) bool {
//      if root.Val == target && root.Left == nil && root.Right == nil {
//         return true
//     }
//     if root == nil {
//         return false
//     }
//     return IsLeafNode(root.Left,target) || IsLeafNode(root.Right,target) 
// }