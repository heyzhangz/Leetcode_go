package question63

/*
	63. 不同路径 II
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
	机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
	现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 */

/*
	思路: 广度 深度。。写个广度的用个队列，超时了。。用DP
	题解:
		用DP 滚动数组优化 只记录上一行的值
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	if len(obstacleGrid) == 0 {
		return 0
	}

	//优化 matrix := make([]int, len(obstacleGrid[0]))
	matrix := make([][]int, len(obstacleGrid))
	for i := range matrix {
		matrix[i] = make([]int, len(obstacleGrid[0]))
	}

	if obstacleGrid[0][0] == 0 {
		matrix[0][0] = 1
	}

	for i, in := range obstacleGrid {
		for j, jn := range in {

			if jn == 1 {
				matrix[i][j] = 0
				continue
			}

			if i - 1 >= 0 {
				matrix[i][j] += matrix[i - 1][j]
			}

			if j - 1 >= 0 {
				matrix[i][j] += matrix[i][j - 1]
			}
		}
	}

	return matrix[len(obstacleGrid) - 1][len(obstacleGrid[0]) - 1]
}