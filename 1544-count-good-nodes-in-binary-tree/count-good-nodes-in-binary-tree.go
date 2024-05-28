/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    count := 0
    s := &MaxStack{}
    s.CountNode(root,&count)
    return count
}

func (s *MaxStack) CountNode(root *TreeNode,count *int) {
    if root == nil {
        return 
    }

    s.Push(root.Val)
    if s.max[len(s.max)-1] <= root.Val {
        *count++
    }
    s.CountNode(root.Left,count)
    s.CountNode(root.Right,count)
    s.Pop()
}

type MaxStack struct {
    stack []int
    max []int
}

func (s *MaxStack) Push(val int) {
    s.stack = append(s.stack,val)
    if len(s.max) <= 0 || s.max[len(s.max)-1] <= val {
        s.max = append(s.max,val)
    }
}

func(s *MaxStack) Pop() {
    if s.stack[len(s.stack)-1] == s.max[len(s.max)-1] {
        s.max = s.max[:len(s.max)-1]
    }
    s.stack = s.stack[:len(s.stack)-1]
}