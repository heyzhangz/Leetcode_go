package question97

/*
	97. 交错字符串
	给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
*/

/*
	思路: DP s3的字符长度取决于s1 s2的几位 即 s3[3] = s1[2]s2[0] + s1"3" || s1[2]s2[0] + s2"1" ...
	题解:
*/

func isInterleave(s1 string, s2 string, s3 string) bool {

	if len(s1) + len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1) + 1)
	for i := range dp {
		dp[i] = make([]bool, len(s2) + 1)
	}

	dp[0][0] = true

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {

			s3loc := i + j - 1
			if i > 0 {
				dp[i][j] = (dp[i - 1][j] && s3[s3loc] == s1[i - 1]) || dp[i][j]
			}
			if j > 0 {
				dp[i][j] = (dp[i][j - 1] && s3[s3loc] == s2[j - 1]) || dp[i][j]
			}
		}
	}

	return dp[len(s1)][len(s2)]
}