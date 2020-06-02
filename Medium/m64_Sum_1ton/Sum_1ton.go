package questionm64

/*
	面试题64. 求1+2+…+n
	求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

	1 <= n <= 10000
 */

/*
	思路: 有1说1 这题真没思路
	题解:
		1. 利用&&逻辑运算符短路的特性 拿递归做
		2. 利用求和公式 快速乘，展开循环即可，最多14位
*/

//1
func sumNums(n int) int {

	ans := 0

	var sum func(int) bool
	sum = func(n int) bool {

		ans += n

		return n > 0 && sum(n - 1)
	}

	sum(n)
	return ans
}

//2 匿名函数
func sunNumsA(n int) int {

	_ = n > 0 && func() bool {
		n += sunNumsA(n - 1)
		return true
	}()

	return n
}

//3 快速乘
func sumNumsQ(n int) int {

	ans, A, B := 0, n, n + 1
	addGreatZero := func() bool {
		ans += A
		return ans > 0
	}

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	return ans >> 1
}