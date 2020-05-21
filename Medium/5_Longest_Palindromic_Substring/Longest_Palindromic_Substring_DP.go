package question5

/*
	5. Longest Palindromic Substring
	Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

	5. 最长回文子串
	给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 */

/*
	题解: 1. 动态规划，没啥意思, time O(n^2) space O(n^2)
	最长回文串的子串也是回文串，利用这个状态进行更新，，即{0, 4}取决于{1, 3}
*/

func longestpalindromeDp(s string) string {

	maxlen := 1
	ans := ""
	statussq := make([][]bool, len(s))
	for i, _ := range s {
		statussq[i] = make([]bool, len(s))
		for j, _ := range s {
			if i == j {
				statussq[i][j] = true
			} else {
				statussq[i][j] = false
			}
		}
	}

	//从头遍历按长度算
	for l := 1; l < len(s); l++ {
		for i := 0; l + i < len(s); i++ {
			j := i + l

			if l == 1 {
				if s[i] == s[j] {
					statussq[i][j] = true
					maxlen = 2
					ans = s[i : j + 1]
				}
			} else if s[i] == s[j] {
				if statussq[i + 1][j - 1] {
					statussq[i][j] = true
					if l + 1 > maxlen {
						maxlen = l + 1
						ans = s[i : j + 1]
					}
				}
			} else {
				statussq[i][j] = false
			}

		}
	}

	return ans
}
