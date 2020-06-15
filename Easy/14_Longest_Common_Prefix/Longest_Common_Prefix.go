package question14

/*
	14. 最长公共前缀

	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""。
 */

/*
	思路: 就纵向地比
	题解: 分治法 两个两个解决，其实也没快
*/

func longestCommonPrefix(strs []string) string {

	if strs == nil || len(strs) == 0 {
		return ""
	}

	minlen := len(strs[0])
	for _, n := range strs {
		if len(n) < minlen {
			minlen = len(n)
		}
	}

	for i := 0; i < minlen; i++ {
		perchar := strs[0][i]
		for _, n := range strs[1:] {
			if n[i] != perchar {
				return n[:i]
			}
		}
	}

	return strs[0][:minlen]
}
