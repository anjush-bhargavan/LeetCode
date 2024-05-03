/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return root
    }
    if key < root.Val {
        root.Left = deleteNode(root.Left,key)
    }else if key > root.Val {
        root.Right = deleteNode(root.Right,key)
    }else{
        if root.Right == nil && root.Left == nil{
            return nil
        }
        if root.Right == nil {
            return root.Left
        }
        if root.Left == nil {
            return root.Right
        }
        root.Val = min(root.Right)
        root.Right = deleteNode(root.Right,root.Val)
    }
    return root
}

func min(node *TreeNode) int {
    if node.Left != nil {
        return min(node.Left)
    }
    return node.Val
}