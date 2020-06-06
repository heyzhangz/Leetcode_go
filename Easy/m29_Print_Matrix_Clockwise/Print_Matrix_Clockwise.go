package questionm29
/*
	面试题29. 顺时针打印矩阵
	输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
 */

/*
	思路: 取长短按顺序来 从外到内，相同的子问题
	题解:
		1. 另开一个二维数组，和顺时针转向数组，通过到头或者重复元素判断
*/

func spiralOrder(matrix [][]int) []int {

	mwide := len(matrix)
	if mwide <= 0 {
		return []int{}
	}
	mlen := len(matrix[0])

	ans := make([]int, mlen * mwide)

	start := 0
	lend := mlen - 1
	wend := mwide - 1

	anscount := 0

	for ; anscount < len(ans); {

		//特殊情况先判断退出，不然影响后面
		if start == lend && start == wend {
			ans[anscount] = matrix[start][start]
			anscount++
			continue
		}

		if start == lend {
			for j := start; j <= wend; j++ {
				ans[anscount] = matrix[j][start]
				anscount++
			}
			continue
		}

		if start == wend {
			for j := start; j <= lend; j++ {
				ans[anscount] = matrix[start][j]
				anscount++
			}
			continue
		}

		//遍历第一个横
		for j := start; j < lend; j++ {
			ans[anscount] = matrix[start][j]
			anscount++
		}

		for j := start; j < wend; j++ {
			ans[anscount] = matrix[j][lend]
			anscount++
		}

		for j := lend; j > start; j-- {
			ans[anscount] = matrix[wend][j]
			anscount++
		}

		for j := wend; j > start; j-- {
			ans[anscount] = matrix[j][start]
			anscount++
		}

		start += 1
		lend -= 1
		wend -= 1
	}

	return ans
}