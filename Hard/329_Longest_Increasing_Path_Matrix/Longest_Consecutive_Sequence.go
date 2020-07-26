package question329

/*
	329. 矩阵中的最长递增路径
	给定一个整数矩阵，找出最长递增路径的长度。
	对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。
*/

/*
	思路: 7月DP月，这题我用DP
	题解:
		拓扑可还行
*/

func longestIncreasingPath(matrix [][]int) int {

	dp := make([][]int, len(matrix))

	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	maxlen := 0

	for i := range matrix {
		for j := range matrix[0] {

			if dp[i][j] == -1 {
				findMaxLen(i, j, &dp, &matrix, &maxlen)
			}
		}
	}

	return maxlen
}

func findMaxLen(x int, y int, dp *[][]int, matrix *[][]int, max *int) int {

	maxlen := 1

	if (*dp)[x][y] != -1 {
		return (*dp)[x][y]
	}

	if x < len(*matrix) - 1 {
		if (*matrix)[x + 1][y] > (*matrix)[x][y] {
			tmplen := 1
			tmplen += findMaxLen(x + 1, y, dp, matrix, max)

			if tmplen > maxlen {
				maxlen = tmplen
			}
		}
	}

	if x > 0 {
		if (*matrix)[x - 1][y] > (*matrix)[x][y] {
			tmplen := 1
			tmplen += findMaxLen(x - 1, y, dp, matrix, max)

			if tmplen > maxlen {
				maxlen = tmplen
			}
		}
	}

	if y < len((*matrix)[0]) - 1 {
		if (*matrix)[x][y + 1] > (*matrix)[x][y] {
			tmplen := 1
			tmplen += findMaxLen(x, y + 1, dp, matrix, max)

			if tmplen > maxlen {
				maxlen = tmplen
			}
		}
	}

	if y > 0 {
		if (*matrix)[x][y - 1] > (*matrix)[x][y] {
			tmplen := 1
			tmplen += findMaxLen(x, y - 1, dp, matrix, max)

			if tmplen > maxlen {
				maxlen = tmplen
			}
		}
	}

	(*dp)[x][y] = maxlen
	if maxlen > *max {
		*max = maxlen
	}
	return maxlen
}