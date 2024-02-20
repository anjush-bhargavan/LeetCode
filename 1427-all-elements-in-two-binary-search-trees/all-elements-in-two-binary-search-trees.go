/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    arr1,arr2 :=[]int{},[]int{}
    inorder(root1,&arr1)
    inorder(root2,&arr2)
    arr := mergeSort(arr1,arr2)
    return arr
}

func inorder(node *TreeNode,arr *[]int) {
    if node == nil {
        return 
    }
    inorder(node.Left,arr)
    *arr = append(*arr,node.Val)
    inorder(node.Right,arr) 
}

func mergeSort(arr1,arr2 []int) []int {
	combined:=[]int{}
	i,j:=0,0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
            combined = append(combined,arr1[i])
            i++
		}else{
			combined = append(combined,arr2[j])
            j++
		}
	}
	for i < len(arr1) {
	    combined = append(combined,arr1[i])
        i++
	}
	for j < len(arr2) {
		combined = append(combined,arr2[j])
        j++
	}
	return combined
}