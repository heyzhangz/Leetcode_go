package question9

/*
	9. 回文数
	判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
	你能不将整数转为字符串来解决这个问题吗？
*/

/*
	思路: 考虑倒序来写数字然后作比较
	题解: 倒序一半，和另一半比较。。题解说是为了避免溢出，但回文数显然不会溢出。。不过能快点
*/

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	oldx := x
	newx := 0
	for oldx != 0 {
		temp := oldx % 10
		oldx /= 10
		newx = newx * 10 + temp
	}

	if x - newx == 0 {
		return true
	}

	return false
}

func isPalindrome_half(x int) bool {

	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber * 10 + x % 10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber / 10
}