package question1431

/*
	1431. 拥有最多糖果的孩子
	给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。
	对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。
*/

/*
	思路: 真就儿童节快乐？
	题解: 枚举可还行
*/

func kidsWithCandies(candies []int, extraCandies int) []bool {

	max := 0
	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}

	ans := make([]bool, len(candies))
	for i, candy := range candies {
		if candy + extraCandies >= max {
			ans[i] = true
		}
	}

	return ans
}