package question198

/*
	198. House Robber

	You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
	Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

 	198. 打家劫舍
	你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
	给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
 */

/*
	思路: 第一个思路是递归，然后光荣超时。。再改成循环，也就是DP
	题解: DP
*/

//改循环 其实就是动规
func rob(nums []int) int {

	if len(nums) <= 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	statue1 := nums[0] //i-2的值
	statue2 := nums[1] //i-1的值
	if nums[0] > nums[1] {
		statue2 = nums[0]
	}
	cur := statue2

	for i := 2; i < len(nums); i++ {
		cur = statue1 + nums[i]
		if cur < statue2 {
			cur = statue2
		}

		statue1 = statue2
		statue2 = cur
	}

	return cur
}

//递归 超时
//func rob(nums []int) int {
//
//	if len(nums) <= 0 {
//		return 0
//	}
//
//	if len(nums) == 1 {
//		return nums[0]
//	}
//
//	num1 := nums[0] + rob(nums[2:])
//	num2 := rob(nums[1:])
//
//	if num1 > num2 {
//		return num1
//	} else {
//		return num2
//	}
//}

