package question108

/*
	108. 将有序数组转换为二叉搜索树
    将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
    本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

*/

/*
	思路: 双指针
	题解: 
        啥破题解！

*/


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

    length := len(nums)

    if length == 0 {
        return nil
    }

    mid := length / 2
    node := TreeNode{Val: nums[mid], Left: nil, Right: nil}
    node.Left = sortedArrayToBST(nums[:mid])
    node.Right = sortedArrayToBST(nums[mid + 1:])

    return &node
}