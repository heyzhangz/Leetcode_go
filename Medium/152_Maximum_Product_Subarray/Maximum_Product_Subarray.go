package question152

/*
	152. 乘积最大子数组
	给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
*/

/*
	思路: 第一反应是DP，仔细一想 应该考虑正负。。
		1. 0 是分隔符，将数组划分成两个子数组，除非只有单个元素 否则不可能是0
		2. 负数同理，只有单个元素才考虑是答案，在算子数组乘积的时候要记录左边第一个负数的总乘积和右边第一个 用来判断哪个更合适截取
		3. 正数子数组最大的就是自己
	题解:
		1. DP 考虑的是正负性的反转，以[i]结尾的最大乘积 取决于 max[i-1] * [i] , min[i-1] * [i], [i]中最大的一个 保留三个状态
			并用滚动数组压缩空间
*/

func calmaxmul(mul int, length int, left int, right int) int {

	if length <= 1 {
		return mul
	}

	if mul > 0 {
		return mul
	}

	if left < right {
		left = right
	}
	mul /= left

	return mul
}

func maxProduct(nums []int) int {

	if len(nums) <= 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	max := nums[0]
	//左右两边最近负数
	left := 1
	right := 1
	submul := 1 //子数组乘积
	start := 0

	for i := 0; i < len(nums); i++ {

		curnum := nums[i]
		if curnum == 0 {
			//到0开始结算
			sublen := i - start//前数组长度
			if sublen > 0 {
				curmax := calmaxmul(submul, sublen, left, right)
				if curmax > max {
					max = curmax
				}
			}

			//算下0自己
			if 0 > max {
				max = 0
			}

			start = i + 1
			left = 1
			right = 1
			submul = 1
		} else if curnum < 0 {
			//负数 看看是否要添加左值 以及要更新右值
			submul *= curnum
			if left > 0 {
				// left 是正数 还没乘到负值 给乘上; 负数表示已经乘到了第一个负值
				left *= curnum
			}
			//right 遇到新的负值就更新
			right = curnum
		} else {
			submul *= curnum
			right *= curnum
			if left > 0 {
				// left 是正数 还没乘到负值
				left *= curnum
			}
		}
	}

	//算最后的结尾
	sublen := len(nums) - start
	if sublen > 0 {
		curmax := calmaxmul(submul, sublen, left, right)
		if curmax > max {
			max = curmax
		}
	}

	return max
}

func maxProduct_DP(nums []int) int {

	maxF, minF, ans := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {

		mx, mn := maxF, minF

		maxF = max(mx * nums[i], max(nums[i], mn * nums[i]))
		minF = min(mn * nums[i], min(nums[i], mx * nums[i]))

		ans = max(maxF, ans)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}