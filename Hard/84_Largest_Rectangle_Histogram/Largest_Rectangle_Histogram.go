package question84

/*
	84. 柱状图中最大的矩形
	给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
	求在该柱状图中，能够勾勒出来的矩形的最大面积。
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