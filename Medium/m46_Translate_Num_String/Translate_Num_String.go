package questionm46

import "strconv"

/*
	面试题46. 把数字翻译成字符串
	给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
 */

/*
	思路: 很显然的DP, 算总数的话 初始化P[n,n] P[n-1, n]...P[n-2, n] = P[n-2] * P[n-1, n] + P[n-2, n-1] * P[n,n]
	题解: 思路一样 滚动数组 + DP，可以写的更简洁。。
*/

func isValid(num int) bool {

	if num >= 0 && num <= 25 {
		return true
	}

	return false
}

func translateNum(num int) int {

	if num < 0 {
		return 0
	}

	if num < 10 {
		return 1
	}

	numstr := strconv.Itoa(num)
	length := len(numstr)
	dp := make([]int, length) //dp[0] 表示 [0, n-1]上可能的结果

	var curnum, prenum int
	//确保至少两位数 初始化后两个
	curnum = num % 10
	num /= 10
	if isValid(curnum) {
		dp[length - 1] = 1
	} else {
		dp[length - 1] = 0
	}
	prenum = curnum
	curnum = num % 10
	num /= 10
	if isValid(curnum) {
		dp[length - 2] += dp[length - 1]
	}
	if isValid(curnum * 10 + prenum) && curnum != 0 {
		dp[length - 2] += 1
	}
	prenum = curnum

	for i := length - 3; i >= 0; i-- {
		curnum = num % 10
		num /= 10
		if isValid(curnum) {
			dp[i] += dp[i + 1]
		}
		if isValid(curnum * 10 + prenum) && curnum != 0 {//哇 贼坑 06不算
			dp[i] += dp[i + 2]
		}
		prenum = curnum
	}

	return dp[0]
}

//题解 直接操作字符串
func translateNum_ans(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := src[i-1:i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}