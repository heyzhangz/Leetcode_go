package question739

/*
	739. 每日温度
	根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。
	例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
	提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
 */

/*
	思路: 傻逼了 这是个栈！
	题解: 也是单调栈
*/

func dailyTemperatures(T []int) []int {

	if len(T) == 0 {
		return []int{}
	}

	//没必要
	ans := make([]int, len(T))
	//for i := range ans {
	//	ans[i] = 0
	//}

	var stack []int

	for i, n := range T {
		if len(stack) == 0 || T[stack[len(stack) - 1]] >= n {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 {
			stacktopindex := stack[len(stack) - 1]
			stacktop := T[stacktopindex]
			if stacktop >= n {
				break
			}
			ans[stacktopindex] = i - stacktopindex
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	return ans
}