package question41

/*
	41. 缺失的第一个正数
	给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
	你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。
*/

/*
	思路: 这。。没啥思路
	题解: 只考虑1-N的数，所以自建hash表，算一算
*/

func firstMissingPositive(nums []int) int {

	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num - 1] = -abs(nums[num - 1])
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}