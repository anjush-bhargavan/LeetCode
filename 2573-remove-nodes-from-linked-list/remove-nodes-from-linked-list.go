/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
    arr := []int{}
    curr := head 
    for curr != nil {
        for len(arr) != 0 && arr[len(arr)-1] < curr.Val {
            arr = arr[:len(arr)-1]
        }
        if len(arr) == 0 || arr[len(arr)-1] >= curr.Val {
            arr = append(arr,curr.Val)
        }
        curr = curr.Next
    }

   newList := &ListNode{Val:arr[0]}
   curr = newList
   for i := 1 ; i < len(arr) ; i++ {
       curr.Next = &ListNode{Val:arr[i]}
       curr = curr.Next
   }
    return newList
}