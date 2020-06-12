package question70

/*
	70. 爬楼梯
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	注意：给定 n 是一个正整数

*/

/*
	思路: 这做了好多次，递归或者滚动数组(DP)，注意滚动组 n 的状态取决于 n-1 和 n-2。。所以只要记录前两个和就行，注意1的边界
		 递归没啥难度 注意递归计算了好多重复状态！哈蛤 果然超时了
	题解: 快速幂 斐波那契公式

*/

func climbStairs(n int) int {

	if n <= 1 {
		return 1
	}

	n1, n2 := 2, 1 //n1 表示 n-1 n2 表示 n - 2

	for i := 3; i <= n; i++ {
		cur := n2 + n1
		n2 = n1
		n1 = cur
	}

	return n1
}

func climbStairs_re(n int) int {

	if n <= 1 {
		return 1
	}

	return climbStairs_re(n - 1) + climbStairs_re(n - 2)
}