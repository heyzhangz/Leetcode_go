package question84

//单调栈题解
func largestrectangleareaS(heights []int) int {

	n := len(heights)
	left, right := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		right[i] = n
	}

	var monoStack []int

	for i := 0; i < n; i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			right[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i] - left[i] - 1) * heights[i])
	}
	return ans
}

func max(x, y int) int {

	if x > y {
		return x
	}

	return y
}