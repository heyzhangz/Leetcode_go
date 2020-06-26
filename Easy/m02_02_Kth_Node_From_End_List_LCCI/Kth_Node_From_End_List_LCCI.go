package questionm02_02
/*
	面试题 02.02. 返回倒数第 k 个节点
	实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
 */

/*
	思路: 俩指针玩
	题解:
*/


type ListNode struct {
	Val int
    Next *ListNode
}

func kthToLast(head *ListNode, k int) int {

    target := head
    for k > 0 {
        head = head.Next
        k--        
    }

    for head != nil {
        head = head.Next
        target = target.Next
    }

    return target.Val
}