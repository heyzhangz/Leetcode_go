package question112

/*
	112. 路径总和
	给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
 */

/*
	思路: 递归迭代都行。。
	题解:
		没差
*/


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
    
    if root == nil {
        return false
    }

    if sum == root.Val && root.Left == nil && root.Right == nil {
        return true
    }
    
    ans := false

    if root.Left != nil {
        ans = ans || hasPathSum(root.Left, sum - root.Val)
    }

    if root.Right != nil {
        ans = ans || hasPathSum(root.Right, sum - root.Val)
    }

    return ans
}

