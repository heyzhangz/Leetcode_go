package question343

/*
	343. 整数拆分
	给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
 */

/*
	思路: 无他，DP
	题解:
		很好，是数学题
*/

func integerBreak(n int) int {

	dp := make([]int, n + 1)
	dp[1] = 1

	for i := 2; i <= n; i++ {
		max := 1
		for j := 1; j < i; j++ {
			cur := j * returnMax(i - j, dp[i - j])
			if cur > max {
				max = cur
			}
		}
		dp[i] = max
	}

	return dp[n]
}

func returnMax(j int, dpj int) int {
	if j > dpj {
		return j
	}

	return dpj
}