package question665

import "math"

/*
	665. 非递减数列
	给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
	我们是这样定义一个非递减数列的： 对于数组中所有的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

*/

/*
	思路: 这题最开始的思路都有小问题，其实还是得考虑特殊情况
	题解:
		从头和尾分别找第一个逆序点，如果相差一个数ok，且大小满足逆序ok
		https://blog.csdn.net/qq_16151925/article/details/80204834
*/

func checkPossibility(nums []int) bool {

	i := 0
	j := len(nums) - 1

	for i < j && nums[i] <= nums[i+1] {
		i++
	}

	for i < j && nums[j] >= nums[j-1] {
		j--
	}

	// 处理首尾边界问题
	head := math.MinInt32
	if i != 0 {
		head = nums[i - 1]
	}

	tail := math.MaxInt32
	if j != len(nums) - 1 {
		tail = nums[j + 1]
	}

	if j - i <= 1 && (head <= nums[j] || nums[i] <= tail) {
		return true
	}

	return false
}