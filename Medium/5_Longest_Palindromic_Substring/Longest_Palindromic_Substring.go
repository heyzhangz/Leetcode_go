package question5

/*
	5. Longest Palindromic Substring Medium
	Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

	5. 最长回文子串
	给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 */

/*
	思路:	从头变量每个字符，查找前后是否有回文；如果向后相邻也相等(偶数情况)，再判断一次 time O(n^2) space O(1)
	题解:
		1. 动态规划，没啥意思, time O(n^2) space O(n^2) 长度为[i,j]的字符串取决于[i+1, j-1]
		2. Manacher算法, time O(n) space O(n)
*/

// optimization1, 改传引用降低内存消耗
func findRalStr(arr *string, findex int, lindex int) int {

	nowlen := 0

	for i := 1; i < len(*arr); i++ {

		if findex - i < 0 || lindex + i >= len(*arr) {
			break
		}

		if (*arr)[findex - i] == (*arr)[lindex + i] {
			nowlen += 1
		} else {
			break
		}
	}

	return nowlen
}

func longestPalindrome(s string) string {

	//chararr := []byte(s) // warning1 golang字符串直接操作
	maxlen := 1
	startindex := 0
	lastindex := 0
	exlen := 0

	// error1 边界错误处理
	if len(s) <= 0 {
		return ""
	}

	for i := 0; i < len(s); i++  {
		nowexlen := findRalStr(&s, i, i)
		nowlen := 1 + 2 * nowexlen
		if nowlen > maxlen {
			maxlen = nowlen
			startindex = i
			lastindex = i
			exlen = nowexlen
		}

		if i + 1 < len(s) && s[i] == s[i + 1] {
			nowexlen := findRalStr(&s, i, i + 1)
			nowlen := 2 + 2 * nowexlen
			if nowlen > maxlen {
				maxlen = nowlen
				startindex = i
				lastindex = i + 1
				exlen = nowexlen
			}
		}
	}

	return s[startindex - exlen : lastindex + exlen + 1]
}
