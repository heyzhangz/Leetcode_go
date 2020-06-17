package question105

/*
	1014. 最佳观光组合
	给定正整数数组 A，A[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的距离为 j - i。
	一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i - j）：景点的评分之和减去它们两者之间的距离。
	返回一对观光景点能取得的最高分。
*/

/*
	思路: 没啥好思路，直接枚举肯定超时，因为最终结果肯定是个正数，所以用j - i <= A[i] + A[j]做了个剪枝，如果后面存在一个大数的话，那最大值也应该是相邻的因为i到j已经抵消掉了
	题解:
		非常tmd妙！将最高分计算拆成A[i] + i 和 A[j] - j，即在[i + 1, max]中找到最大的 A[j] - j 和 之前最大的 A[i] + i
*/


func maxScoreSightseeingPair(A []int) int {

	if len(A) <= 1 {
		return 0
	}

	max := 0
	for i := 0; i < len(A) - 1; i++ {
		for j := i + 1; j < len(A) && j - i <= A[i] + A[j]; j++ {
			cur := A[i] + A[j] + i - j
			if cur > max {
				max = cur
			}
		}
	}

	return max
}

func maxScoreSightseeingPair_ans(A []int) int {
	ans, mx := 0, A[0] + 0
	for j := 1; j < len(A); j++ {
		ans = max(ans, mx + A[j] - j)
		// 边遍历边维护
		mx = max(mx, A[j] + j)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}