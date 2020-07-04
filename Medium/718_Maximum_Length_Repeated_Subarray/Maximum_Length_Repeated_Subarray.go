package question718

/*
	718. 最长重复子数组
	给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
 */

/*
	思路: 想到是DP，没想到咋记录。。偷偷看了眼提示
	题解:
		DP 可以用滚动数组优化空间
*/

func findLength(A []int, B []int) int {

	maxlen := 0
	matrix := make([][]int, len(A) + 1)
	for i := range matrix {
		matrix[i] = make([]int, len(B) + 1)
	}

	for i := len(A) - 1; i >= 0; i-- {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				matrix[i][j] = matrix[i + 1][j + 1] + 1
				if matrix[i][j] > maxlen {
					maxlen = matrix[i][j]
				}
			} else {
				matrix[i][j] = 0 // 最大前缀
			}
		}
	}

	return maxlen
}