package questionm02_05
/*
	面试题 02.05. 链表求和
	给定两个用链表表示的整数，每个节点包含一个数位。
    这些数位是反向存放的，也就是个位排在链表首部。
    编写函数对这两个整数求和，并用链表形式返回结果。
 */

/*
	思路: 就加
	题解:
*/


type ListNode struct {
	Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    pre := 0
    var ans, tail *ListNode

    for l1 != nil || l2 != nil {

        a := 0
        b := 0
        now := 0
        if l1 != nil {
            a = l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            b = l2.Val
            l2 = l2.Next
        }

        now = a + b + pre
        pre = now / 10
        now %= 10

        if tail == nil {
            tail = &ListNode{Val: now, Next: nil}
            ans = tail
        } else {
            tail.Next = &ListNode{Val: now, Next: nil}
            tail = tail.Next
        }
    }

    if pre != 0 && tail != nil{
        tail.Next = &ListNode{Val: pre, Next: nil}
        tail = tail.Next
    }

    return ans
}