/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 var rem int

func doubleIt(head *ListNode) *ListNode {
    rem = 0
    head = add(head)
    if rem != 0 {
        
        newNode := &ListNode{Val:rem,Next:head}
        return newNode
    }
    return head
}

func add(head *ListNode) *ListNode{
    if head.Next == nil {
        n := head.Val * 2
        if n > 9 {
            head.Val = n%10
            rem =  n/10
        }else{
            head.Val = n
        }
        return head
    }
    head.Next = add(head.Next)
    n := (head.Val * 2) + rem
    if n > 9 {
        head.Val = n%10
        rem = n/10
    }else{
        head.Val = n
        rem = 0
    } 
    return head
}