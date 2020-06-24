package question16

import "sort"

/*
	15. 三数之和
	给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 */

/*
	思路: 范围不大，暴力？暴力A了
	题解:
		这个还是双指针
*/

func abs(n int) int {
	y := n >> 31
	return (n ^ y) - y
}

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	ans := nums[0] + nums[1] + nums[2]
	ansc := abs(ans - target)

	for i := range nums {
		j := i + 1
		k := len(nums) - 1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < k {
			cur := nums[i] + nums[j] + nums[k]
			curc := cur - target

			if curc == 0 {
				return cur
			} else if curc > 0 {
				if curc < ansc {
					ans = cur
					ansc = curc
				}
				k--
			} else {
				if -curc < ansc {
					ans = cur
					ansc = -curc
				}
				j++
			}
		}
	}

	return ans
}