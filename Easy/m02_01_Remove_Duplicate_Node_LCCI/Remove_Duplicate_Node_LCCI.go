package questionm02_01
/*
	面试题 02.01. 移除重复节点
	编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
 */

/*
	思路: 这。。hashmap
	题解:
*/


type ListNode struct {
	Val int
    Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {

    if head == nil {
        return nil
    }

    a := make(map[int]bool)
    a[head.Val] = true
    ans := head
    pre := head
    head = pre.Next

    for head != nil {

        if a[head.Val] {
            pre.Next = head.Next
        } else {
            pre = head
            a[head.Val] = true
        }
        head = pre.Next
    }

    return ans
}