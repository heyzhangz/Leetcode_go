package question1300

import "sort"

/*
	1300. 转变数组后最接近目标值的数组和
	给你一个整数数组 arr 和一个目标值 target ，请你返回一个整数 value ，使得将数组中所有大于 value 的值变成 value 后，数组的和最接近  target （最接近表示两者之差的绝对值最小）。
	如果有多种使得和最接近 target 的方案，请你返回这些整数中的最小值。
	请注意，答案不一定是 arr 中的数字。
 */

/*
	思路: 没啥思路，看了看提示竟然用二分查找；最大值就是数组最大值，最小值是平均值，然后在中间找合理的
	题解:
		题解还不如我的，就是用了个排序前缀和加速了下求和
		倒是下面有个有意思的题解 先按递增排序；然后一趟遍历，发现当下值大于（目标值减去累加和除以剩下个数的值）时，则返回剩余的平均值，思路不难理解吧？

*/

func abs(n int) int {
	y := n >> 31
	return (n ^ y) - y
}

func BinSearch(start int, end int, arr []int, target int) int {

	startsum := 0
	endsum := 0

	if start >= end {
		return end //一定要返回end，max来自于数组
	}

	for _, n := range arr {

		if n > start {
			startsum += start
		} else {
			startsum += n
		}

		if n > end {
			endsum += end
		} else {
			endsum += n
		}
	}

	startsum = abs(startsum - target)
	endsum = abs(endsum - target)

	if end - start == 1 || start - end == 1 {
		if startsum > endsum {
			return end
		} else {
			return start
		}
	}

	if startsum > endsum {
		return BinSearch((start + end) / 2, end, arr, target)
	} else {
		return BinSearch(start, (start + end) / 2, arr, target)
	}

}

func findBestValue(arr []int, target int) int {

	min := target / len(arr)
	max := 0
	for _, n := range arr {
		if n > max {
			max = n
		}
	}

	return BinSearch(min, max, arr, target)
}

func findBestValue_dir(arr []int, target int) int {

	if arr == nil || len(arr) == 0 {
		return 0
	}

	sort.Ints(arr)
	cursum := 0

	for i, n := range arr {
		x := (target - cursum) / (len(arr) - i)
		if x < n {
			t := (float32(target) - float32(cursum)) / (float32(len(arr) - i))
			//t := (float32(target) - float32(cursum)) / (float32(len(arr) - i))
			//return int(t + 0.49)
			//四舍五入
			if t - float32(x) > 0.5 {
				return x + 1
			} else {
				return x
			}
		}
		cursum += n
	}

	return arr[len(arr) - 1]
}