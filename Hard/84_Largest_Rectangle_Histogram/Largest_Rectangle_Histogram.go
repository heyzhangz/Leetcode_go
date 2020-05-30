package question84

/*
	76. Minimum Window Substring
	Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

	76. 最小覆盖子串
	给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。

	如果 S 中不存这样的子串，则返回空字符串 ""。
	如果 S 中存在这样的子串，我们保证它是唯一的答案。
*/

/*
	思路: 遍历找哪个比它矮，记状态然后可以直接用
	题解:
*/

func findLess(heights *[]int, start int) (int, int) {

	left := start
	right := start

	for ; left >= 0; left-- {
		if (*heights)[left] < (*heights)[start] {
			break
		}
	}

	for ; right < len(*heights); right++ {
		if (*heights)[right] < (*heights)[start] {
			break
		}
	}

	return left + 1, right - 1
}

func findLessWithLeft(heights *[]int, start int, left int) (int, int) {

	right := start

	for ; left >= 0; left-- {
		if (*heights)[left] < (*heights)[start] {
			break
		}
	}

	for ; right < len(*heights); right++ {
		if (*heights)[right] < (*heights)[start] {
			break
		}
	}

	return left + 1, right - 1
}

func largestRectangleArea(heights []int) int {

	if len(heights) <= 0 {
		return 0
	}

	areamap := make([]int, len(heights))
	for i, _ := range areamap {
		areamap[i] = -1
	}

	for i := 0; i < len(heights); i++ {

		//已经判断过了跳过
		if areamap[i] != -1 {
			continue
		}

		start := i
		//没判断过 向左向右找更矮的
		left, right := findLess(&heights, start)
		areamap[i] = (right - left + 1) * heights[start]
		for ; ; {
			//右边的最小值可以用来判断
			if right >= len(heights) - 1 || areamap[right + 1] != -1 {
				break
			}

			left = start
			start = right + 1
			left, right := findLessWithLeft(&heights, start, left)
			areamap[start] = (right - left + 1) * heights[start]
		}
	}

	max := -1
	for _, a := range areamap {
		if a > max {
			max = a
		}
	}

	return max
}