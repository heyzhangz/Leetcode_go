package question114

/*
	114. 二叉树展开为链表
	给定一个二叉树，原地将它展开为一个单链表。
 */

/*
	思路: 前序遍历展开
	题解:
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {

	if root == nil {
		return
	}

	preOrderFlatten(root)
}

func preOrderFlatten(root *TreeNode) *TreeNode {

	leftend := root
	rightend := root

	if root.Left != nil {
		leftend = preOrderFlatten(root.Left)
	}

	if root.Right != nil {
		rightend = preOrderFlatten(root.Right)
	}

	end := root
	retLeft := root.Left
	retRight := root.Right
	root.Left = nil
	if retLeft != nil {
		root.Right = retLeft
		end = leftend
	}

	if retRight != nil {
		end.Right = retRight
		end = rightend
	}

	return end
}