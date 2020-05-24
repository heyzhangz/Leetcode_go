package question76

/*
	76. Minimum Window Substring
	Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

	76. 最小覆盖子串
	给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。

	如果 S 中不存这样的子串，则返回空字符串 ""。
	如果 S 中存在这样的子串，我们保证它是唯一的答案。
*/

/*
	思路: 动态滑动窗口
	题解:
*/

//这回真滑，从左到右利用前面的状态
func minWindow(s string, t string) string {

	res := ""
	minlen := len(s)

	if len(t) > len(s) {
		return res
	}

	tmap := make(map[byte]int, len(t))
	sumtmap := 0 //不用判断整个map了
	for i := range t {
		if _, ok := tmap[t[i]]; ok {
			tmap[t[i]] += 1
			sumtmap += 1
		} else {
			tmap[t[i]] = 1
			sumtmap += 1
		}
	}

	//窗口[i,j] 又指针j是从头遍历到尾的
	i := 0
	for j := 0; j < len(s); j++ {
		//从右边加一个值
		if _, ok := tmap[s[j]]; ok {
			tmap[s[j]] -= 1
			if tmap[s[j]] >= 0 {
				sumtmap -= 1
			}
		}
		//判断sum， 满足条件就从左边减一个值
		if sumtmap == 0 {
			for ; i <= j; i++ {
				if _, ok := tmap[s[i]]; ok {
					tmap[s[i]] += 1
					if tmap[s[i]] > 0 {
						sumtmap += 1
						curlen := j - i + 1
						if curlen <= minlen {
							minlen = curlen
							res = s[i : j + 1]
						}
						i++
						break
					}
				}
			}
		}
	}

	return res
}

//func isOK(tmap map[byte]int) bool {
//
//	for key := range tmap {
//		if tmap[key] > 0 {
//			return false
//		}
//	}
//
//	return true
//}

//思路2 ： 你懂个P的滑动窗口 =。=滑动一哈
//func minWindow(s string, t string) string {
//
//	//一个map 统计一下t字符的初始状态
//	tchararr := []byte(t)
//	tmap := make(map[byte]int, len(t))
//	for _, c := range tchararr {
//		if _, ok := tmap[c]; ok {
//			tmap[c] += 1
//		} else {
//			tmap[c] = 1
//		}
//	}
//
//	//记录字串状态，从t的长度开始
//	startlen := len(t)
//	if startlen > len(s) {
//		return ""
//	}
//
//	//开始滑
//	curmap := make(map[byte]int, len(t))
//	//长case超时，记一次上一个状态，每次不用重新遍历map
//	tmapsum := 0
//	for key := range tmap {
//		tmapsum += tmap[key]
//	}
//	for wlen := startlen; wlen <= len(s); wlen++ {
//		//初始化当前状态
//		for key := range tmap {
//			curmap[key] = tmap[key]
//		}
//		curmapsum := tmapsum
//		for _, c := range []byte(s[0 : wlen]) {
//			if _, ok := tmap[c]; ok {
//				curmap[c] -= 1
//				if curmap[c] >= 0 {
//					curmapsum -= 1
//				}
//			}
//		}
//
//		if curmapsum <= 0 {
//			return s[0 : wlen]
//		}
//
//		for i := 0; i <= len(s) - wlen; i++ {
//			if i - 1 < 0 || i + wlen > len(s) {
//				continue
//			}
//			if _, ok := curmap[s[i + wlen - 1]]; ok {
//				curmap[s[i + wlen - 1]] -= 1
//				if curmap[s[i + wlen - 1]] >= 0 {
//					curmapsum -= 1
//				}
//			}
//			if _, ok := curmap[s[i - 1]]; ok {
//				if curmap[s[i - 1]] >= 0 {
//					curmapsum += 1
//				}
//				curmap[s[i - 1]] += 1
//			}
//			if curmapsum <= 0 {
//				return s[i : i + wlen]
//			}
//		}
//	}
//
//	return ""
//}

//思路1 ： 超时 傻逼了。。这哪是动规，这就是个暴力
//func containAndRemove(s byte, tmap *map[byte]int) bool {
//
//	for key := range *tmap {
//		if s == key && (*tmap)[key] > 0{
//			(*tmap)[key] -= 1
//			return true
//		}
//	}
//
//	return false
//}
//
//func isOK(tmap map[byte]int) bool {
//
//	for key := range tmap {
//		if tmap[key] > 0 {
//			return false
//		}
//	}
//
//	return true
//}

//func minWindow(s string, t string) string {
//
//	tchararr := []byte(t)
//	tmap := make(map[byte]int, len(tchararr))
//	for _, c := range tchararr {
//		if _, ok := tmap[c]; !ok {
//			tmap[c] = 1
//		} else {
//			tmap[c]++
//		}
//	}
//
//	statusseq := make([][]map[byte]int, len(s))
//	for i := 0; i < len(s); i++ {
//		statusseq[i] = make([]map[byte]int, len(s))
//		for j := 0; j < len(s); j++ {
//			statusseq[i][j] = make(map[byte]int, len(tchararr))
//			for key := range tmap {
//				statusseq[i][j][key] = tmap[key]
//			}
//			//statusseq[i][j] = tmap
//			if j == 0 {
//				containAndRemove(s[i], &statusseq[i][j])
//			}
//		}
//	}
//
//	for i := 0; i < len(s); i++ {
//		for l := 1; l < len(s) - i; l++ {
//			//statusseq[i][l] = statusseq[i][l - 1]
//			for key := range tmap {
//				statusseq[i][l][key] = statusseq[i][l - 1][key]
//			}
//			containAndRemove(s[i + l], &statusseq[i][l])
//		}
//	}
//
//	minlen := len(s)
//	res := ""
//	for i := 0; i < len(s); i++ {
//		for j := 0; j < len(s); j++ {
//			if isOK(statusseq[i][j]) && j + 1 <= minlen {
//				minlen = j + 1
//				res = s[i : i + j + 1]
//			}
//		}
//	}
//
//	return res
//}