/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    newTree := &TreeNode{Val : preorder[0]}
    for i := 1 ; i < len(preorder) ; i++ {
        Insert(newTree,preorder[i])
    }
    return newTree
}

func Insert(root *TreeNode,data int) *TreeNode {
    if root == nil {
        return &TreeNode{Val:data}
    }
    if data < root.Val {
        root.Left = Insert(root.Left,data)
    }else if data > root.Val {
        root.Right = Insert(root.Right,data)
    }
    return root
}