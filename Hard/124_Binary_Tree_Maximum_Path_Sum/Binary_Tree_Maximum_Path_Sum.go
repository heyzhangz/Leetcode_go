package question124

import "math"

/*
	124. 二叉树中的最大路径和
	给定一个非空二叉树，返回其最大路径和。
	本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
*/

/*
	思路: hard题就这？就这？ 递个归
	题解:
		用闭包很俏皮，之后需要全局量的时候可以考虑这个方法
		还是我的快 应该是变量的原因，我这边是入参
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func getMaxSum(root *TreeNode, allmax *int) int {

	curmax := 0
	sidemax := 0
	if root.Left != nil {
		t := getMaxSum(root.Left, allmax)
		if t > 0 {
			curmax += t
			sidemax = t
		}
	}

	if root.Right != nil {
		t := getMaxSum(root.Right, allmax)
		if t > 0 {
			curmax += t
			if t > sidemax {
				sidemax = t
			}
		}
	}

	t := root.Val + curmax
	sidemax += root.Val
	if t > *allmax {
		*allmax = t
	}

	if sidemax < 0 {
		return 0
	}

	return sidemax
}

func maxPathSum(root *TreeNode) int {

	var ans int = -2147483648
	getMaxSum(root, &ans)

	return ans
}

//非题解 用闭包写一个
func maxPathSum_ans(root *TreeNode) int {

	ans := math.MinInt32

	//important
	var getMax func(root *TreeNode) int
	getMax = func(root *TreeNode) int {

		if root == nil {
			return 0
		}

		leftmax := max(getMax(root.Left), 0)
		rightmax := max(getMax(root.Right), 0)

		fmax := root.Val + leftmax + rightmax
		if fmax > ans {
			ans = fmax
		}

		return max(root.Val + leftmax, root.Val + rightmax)
	}

	getMax(root)

	return ans
}

func max(a ...int) int {

	max := math.MinInt32
	for _, aa := range a {
		if aa > max {
			max = aa
		}
	}

	return max
}