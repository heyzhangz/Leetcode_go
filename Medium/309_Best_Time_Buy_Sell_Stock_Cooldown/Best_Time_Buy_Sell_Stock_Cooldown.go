package question394

/*
	309. 最佳买卖股票时机含冷冻期

	给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
	设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
	你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
	卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
 */

/*
	思路: DP 竟然可能是负数-，-
	题解:
		3*n数组。表示三种状态
*/

func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	if len(prices) == 2 {
		if prices[1] > prices[0] {
			return prices[1] - prices[0]
		}
		return 0
	}

	dp := make([]int, len(prices) + 2)

	for i := 1; i < len(prices); i++ {
		dploc := i + 2
		max := dp[dploc - 1]
		for j := 0; j < i; j++ {
			cur := dp[dploc - i + j - 2] + prices[i] - prices[j]
			if cur > max{
				max = cur
			}
		}
		dp[dploc] = max
	}

	return dp[len(dp) - 1]
}

func maxProfit_ans(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	f := make([][3]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2] - prices[i])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[n-1][1], f[n-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}