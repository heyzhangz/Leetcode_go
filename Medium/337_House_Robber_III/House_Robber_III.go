package question337

/*
	337. 打家劫舍 III
	在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
	计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
 */

/*
	思路: 记忆化遍历
	题解:
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	okMaxlist := make(map[*TreeNode]int)
	noMaxlist := make(map[*TreeNode]int)

	return max(noMaxValue(root, &okMaxlist, &noMaxlist), okMaxValue(root, &okMaxlist, &noMaxlist))
}

func okMaxValue(root *TreeNode, okMaxList *map[*TreeNode]int, noMaxlist *map[*TreeNode]int) int {

	if root == nil {
		return 0
	}

	if _, ok := (*okMaxList)[root]; !ok {
		(*okMaxList)[root] = root.Val + noMaxValue(root.Right, okMaxList, noMaxlist) + noMaxValue(root.Left, okMaxList, noMaxlist)
	}

	return (*okMaxList)[root]
}

func noMaxValue(root *TreeNode, okMaxList *map[*TreeNode]int, noMaxlist *map[*TreeNode]int) int {

	if root == nil {
		return 0
	}

	if _, ok := (*noMaxlist)[root]; !ok {
		(*noMaxlist)[root] = max(okMaxValue(root.Right, okMaxList, noMaxlist), noMaxValue(root.Right, okMaxList, noMaxlist)) +
			max(okMaxValue(root.Left, okMaxList, noMaxlist), noMaxValue(root.Left, okMaxList, noMaxlist))
	}

	return (*noMaxlist)[root]
}

func max(a int, b int) int {

	if a > b {
		return a
	}

	return b
}