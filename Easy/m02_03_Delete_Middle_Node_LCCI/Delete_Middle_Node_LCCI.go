package questionm02_03
/*
	面试题 02.03. 删除中间节点
	实现一种算法，删除单向链表中间的某个节点（即不是第一个或最后一个节点），假定你只能访问该节点。
 */

/*
	思路: 那就生挪呗
	题解:
*/


type ListNode struct {
	Val int
    Next *ListNode
}

func deleteNode(node *ListNode) {

    pre := node
    for node.Next != nil {
        node.Val = node.Next.Val
        pre = node
        node = node.Next
    }

    pre.Next = nil
}