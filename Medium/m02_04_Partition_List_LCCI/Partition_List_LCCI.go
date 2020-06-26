package questionm02_04
/*
	面试题 02.04. 分割链表
	编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。
 */

/*
	思路: 新建两个
	题解:
*/


type ListNode struct {
	Val int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {

    if head == nil {
        return head
    }

    var less, more, lesshead, morehead *ListNode = nil, nil, nil, nil

    for head != nil {
        if head.Val >= x {
            t := &ListNode{Val: head.Val, Next: nil}
            if more == nil {
                more = t
                morehead = more
            } else {
                more.Next = t
                more = t
            }
        } else {
            t := &ListNode{Val: head.Val, Next: nil}
            if less == nil {
                less = t
                lesshead = less
            } else {
                less.Next = t
                less = t
            }
        }
        head = head.Next
    }

    if less != nil {
        less.Next = morehead
        return lesshead
    } else {
        return morehead
    }
}