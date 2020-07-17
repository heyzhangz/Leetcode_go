package question35

/*
	35. 搜索插入位置
	给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
    你可以假设数组中无重复元素。
 */

/*
	思路: 这题一看就是二分，注意边界条件
	题解:
*/


func searchInsert(nums []int, target int) int {

    left := 0
    right := len(nums) - 1

    for left <= right {

        mid := (left + right) >> 1

        if target == nums[mid] {
            return mid
        } else if target < nums[mid] {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    // 这里最后情况是right == left，那如果target > left要返回 + 1 , 否则就是left 这里和取left的逻辑一致
    return left
}
