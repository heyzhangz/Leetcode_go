package question139

/*
	139. 单词拆分
	给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

	说明：
		拆分时可以重复使用字典中的单词。
		你可以假设字典中没有重复的单词。
 */

/*
	思路: 递归 找个一个词之后看子数组
	题解:
		加个记忆 不然重复计算不过
		题解 dp
*/

func contain(s string, wordDict *[]string) bool {

	for _, n := range *wordDict {
		if n == s {
			return true
		}
	}

	return false
}

func isBreak(s string, wordDict *map[string]bool, maxlen int, start int, cache *map[int]bool) bool {

	if _, ok := (*cache)[start]; ok {
		return (*cache)[start]
	}

	_, ok := (*wordDict)[s]
	if start == len(s) || ok {
		(*cache)[start] = true
		return true
	}

	for i := start; i <= len(s); i++ {
		if i - start > maxlen + 1 {
			break
		}
		sub := s[start : i]
		_, ok = (*wordDict)[sub]
		if ok && isBreak(s, wordDict, maxlen, i, cache) {
			(*cache)[start] = true
			return true
		}
	}

	(*cache)[start] = false
	return false
}

func wordBreak(s string, wordDict []string) bool {

	maxlen := 0
	wordmap := make(map[string]bool)
	for _, n := range wordDict {
		if len(n) > maxlen {
			maxlen = len(n)
		}
		wordmap[n] = true
	}

	return isBreak(s, &wordmap, maxlen, 0, &map[int]bool{})
}