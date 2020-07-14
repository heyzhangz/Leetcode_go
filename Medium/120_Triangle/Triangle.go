package question105

import "math"

/*
	120. 三角形最小路径和

	给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
	相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
*/

/*
	思路: DP
	题解:
		我们从 ii 到 00 递减地枚举 jj，这样我们只需要一个长度为 nn 的一维数组 ff，就可以完成状态转移。
*/


func minimumTotal(triangle [][]int) int {

	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		tmp := make([]int, len(triangle[i]))
		tmp[0] = dp[0] + triangle[i][0]
		tmp[len(triangle[i]) - 1] = dp[len(triangle[i]) - 2] + triangle[i][len(triangle[i]) - 1]
		for j := 1; j < len(triangle[i]) - 1; j++ {
			tmp[j] = min(dp[j - 1], dp[j]) + triangle[i][j]
		}

		for i, n := range tmp {
			dp[i] = n
		}
	}

	ans := math.MaxInt32
	for _, n := range dp {
		if n < ans {
			ans = n
		}
	}

	return ans
}

func min(a int, b int) int {

	if a < b {
		return a
	}

	return b
}