package question315

/*
	315. 计算右侧小于当前元素的个数
	给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。
*/

/*
	思路: 这暴力肯定不行，降低复杂度首选二分，从后往前做个排序
	题解:
		BST

*/

func countSmaller(nums []int) []int {

	sortnums := make([]int, len(nums))
	sortlen := 0

	ans := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {

		insertloc := divSearch(sortnums[:sortlen], nums[i])
		ans[i] = insertloc

		sortlen++
		insertnum := nums[i]
		for j := insertloc; j < sortlen; j++ {
			tmp := sortnums[j]
			sortnums[j] = insertnum
			insertnum = tmp
		}
	}

	return ans
}

func divSearch(sortnums []int, a int) int {

	if len(sortnums) == 0 {
		return 0
	}

	left := 0
	right := len(sortnums) - 1

	for right - left > 1 {
		mid := (right + left) / 2

		if a > sortnums[mid] {
			left = mid
			continue
		} else {
			right = mid
			continue
		}
	}

	if a <= sortnums[left] {
		return left
	} else if a <= sortnums[right]{
		return right
	} else {
		return right + 1
	}
}