package question287

/*
	287. Find the Duplicate Number
	Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

	Note:
	You must not modify the array (assume the array is read only).
	You must use only constant, O(1) extra space.
	Your runtime complexity should be less than O(n2).
	There is only one duplicate number in the array, but it could be repeated more than once.

	287. 寻找重复数
	给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

	说明：
	不能更改原数组（假设数组是只读的）。
	只能使用额外的 O(1) 的空间。
	时间复杂度小于 O(n^2) 。
	数组中只有一个重复的数字，但它可能不止重复出现一次。

 */

/*
	思路: 这题的条件很多，时间复杂度小于O(n^2)，先考虑二分，因为最大的数是n，可以根据n/2的左右两边数量来缩小具体位置情况
	题解:
		1. 二进制逐位推测
		2. 转换成一个循环链判断
*/

func findDuplicate(nums []int) int {

	start := 1
	end := len(nums)

	for length := end - start + 1; length > 0; {
		mid := (end - start + 1) / 2 + start
		highnums := 0
		lownums := 0
		midcount := 0
		for _, n := range nums {
			if n < start || n > end {
				continue
			}

			if n > mid {
				highnums++
			} else if n < mid {
				lownums++
			} else {
				midcount++
				if midcount > 1 {
					return mid
				}
			}
		}

		if length % 2 == 0 {
			highnums--
		}

		if lownums > highnums {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return 0
}
