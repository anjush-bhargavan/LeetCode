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
        arr = append(arr,curr.Val)
        curr = curr.Next
    }
    max := arr[len(arr)-1]
    for i := len(arr)-2 ; i >= 0 ; i-- {
        if max > arr[i] {
            arr = append(arr[:i],arr[i+1:]...)
        }
        if max < arr[i] {
            max = arr[i]
        }
    }

   newList := &ListNode{Val:arr[0]}
   curr = newList
   for i := 1 ; i < len(arr) ; i++ {
       curr.Next = &ListNode{Val:arr[i]}
       curr = curr.Next
   }
    return newList
}