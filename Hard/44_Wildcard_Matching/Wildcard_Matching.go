package question44

/*
	44. 通配符匹配
	给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
*/

/*
	思路: 之前10题做过类似的，这回DP试一下
	题解:
		除DP外还有个贪心，即拿*算分隔，匹配字串
*/

func isMatch(s string, p string) bool {

	martix := make([][]bool, len(s) + 1)
	for i := range martix {
		martix[i] = make([]bool, len(p) + 1)
	}

	//初始化
	//dp[0][j]dp[0][j] 需要分情况讨论：因为星号才能匹配空字符串，所以只有当模式 pp 的前 jj 个字符均为星号时，dp[0][j]dp[0][j] 才为真。
	for i := 1; i <= len(p); i++ {
		if p[i - 1] == '*' {
			martix[0][i] = true
		} else {
			break
		}
	}

	martix[0][0] = true

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j - 1] == '*' {
				martix[i][j] = martix[i - 1][j] || martix[i][j - 1]
			} else {
				martix[i][j] = martix[i - 1][j - 1] && (s[i - 1] == p[j - 1] || p[j - 1] == '?')
			}
		}
	}

	return martix[len(s)][len(p)]
}