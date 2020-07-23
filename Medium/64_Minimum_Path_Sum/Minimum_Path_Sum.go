package question64

/*
	64. 最小路径和
	给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
	说明：每次只能向下或者向右移动一步。
 */

/*
	思路: 肯定dp啊
	题解: 思路一样
*/

func minPathSum(grid [][]int) int {

	dp := make([]int, len(grid[0]))
	dp[0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[i] = dp[i - 1] + grid[0][i]
	}

	for _, line := range grid[1:] {

		dp[0] = dp[0] + line[0]
		for i := 1; i < len(line); i++ {
			if dp[i - 1] < dp[i] {
				dp[i] = dp[i - 1] + line[i]
			} else {
				dp[i] += line[i]
			}
		}
	}

	return dp[len(grid[0]) - 1]
}