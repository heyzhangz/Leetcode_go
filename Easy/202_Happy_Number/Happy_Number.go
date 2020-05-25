package question202

/*
	202. Happy Number

	Write an algorithm to determine if a number n is "happy".
	A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.
	Return True if n is a happy number, and False if not.

 	202. 快乐数
	编写一个算法来判断一个数 n 是不是快乐数。
	「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。
	如果 n 是快乐数就返回 True ；不是，则返回 False 。
 */

/*
	思路: 将每次的结果存到一个数组里，如果出现重复的结果 则说明是循环。。优化的话一是数组排序 二是反正就10个数的平方和用个map存
	题解: 数据结构用hashset 即map存
		 用快慢指针，是循环的话两个数总有相等的时候！！
*/

var SQUAREMAP = map[int]int {
	0 : 0,
	1 : 1,
	2 : 4,
	3 : 9,
	4 : 16,
	5 : 25,
	6 : 36,
	7 : 49,
	8 : 64,
	9 : 81,
}

func getBitSum(n int) int {

	sum := 0

	for i := 10; n > 0; {
		sum += SQUAREMAP[(n % i)]
		n /= i
	}

	return sum
}

func contain(n int, ints []int) bool {

	for _, i := range ints {
		if n == i {
			return true
		}
	}

	return false
}

func isHappy(n int) bool {

	sum := 0
	sumarr := map[int]int{}
	for ; sum != 1; {
		sum = getBitSum(n)
		if _, ok := sumarr[sum]; ok {
			return false
		}
		n = sum
		sumarr[sum] = sum
	}
	return true
}

