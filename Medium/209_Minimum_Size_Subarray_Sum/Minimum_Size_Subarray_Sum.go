package question209

import "math"

/*
	209. 长度最小的子数组
	给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。
 */

/*
	思路: 滑动窗口 题目中还要个nlogn的算法，二分
	题解:
*/

func minSubArrayLen(s int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	maxlen := math.MaxInt32
	start, end := 0, 0
	now := nums[start]

	for true {

		if now >= s {
			nowlen := end - start + 1
			if nowlen < maxlen {
				maxlen = nowlen
			}
			now -= nums[start]
			start++
		} else {
			if end == len(nums) - 1 {
				break
			}
			end++
			now += nums[end]
		}
	}

	if maxlen == math.MaxInt32 {
		return 0
	}
	return maxlen
}