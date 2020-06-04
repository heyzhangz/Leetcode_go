package question238
/*
	238. 除自身以外数组的乘积
	给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
 */

/*
	思路: 思路和题解一致，等于左边成右边 但是没从O(2n)的角度上来想
	题解: 1. 左右乘积列表 遍历两遍 空间改O(1)注意最后其实只用相近状态即可，其中一个改常量
		 2. 更极端一点 两个循环其实都是类似的，改双指针
*/

func productExceptSelf(nums []int) []int {

	if len(nums) <= 0 {
		return []int{}
	}

	leftmultis := make([]int, len(nums))
	rightmultis := make([]int, len(nums))

	leftmultis[0] = 1
	rightmultis[len(nums) - 1] = 1

	//左边值 = 左边现有和乘现有值
	for i := 1; i < len(nums); i++ {

		leftmultis[i] = leftmultis[i - 1] * nums[i - 1]
	}

	//右边值同理
	for i := len(nums) - 2; i >= 0; i-- {
		rightmultis[i] = rightmultis[i + 1] * nums[i + 1]
	}

	//空间改O1
	//tempright := 1
	//for i := len(nums) - 2; i >= 0; i-- {
	//	tempright *= nums[i + 1]
	//	leftmultis[i] = tempright * leftmultis[i]
	//}

	//把一个改成ans
	for i, n := range leftmultis {
		leftmultis[i] = n * rightmultis[i]
	}

	return leftmultis
}

func productExceptSelf_doublepoint(nums []int) []int {

	if len(nums) <= 0 {
		return []int{}
	}

	ans := make([]int, len(nums))
	for i := range ans{
		ans[i] = 1
	}

	//左边值 = 左边现有和乘现有值
	templeft := 1
	tempright := 1
	for i, j := 1, len(nums) - 2; i < len(nums); {

		tempright = tempright * nums[j + 1]
		templeft = templeft * nums[i - 1]
		ans[i] *= templeft
		ans[j] *= tempright
		i++
		j--
	}

	return ans
}