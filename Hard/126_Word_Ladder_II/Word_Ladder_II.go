package question126

/*
	126. 单词接龙 II
	给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：
	每次转换只能改变一个字母。
	转换过程中的中间单词必须是字典中的单词。
	说明:
		如果不存在这样的转换序列，返回一个空列表。
		所有单词具有相同的长度。
		所有单词只由小写字母组成。
		字典中不存在重复的单词。
		你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
*/

/*
	思路: 类似于华为的思路，不过是广度有限遍历，得搞个队列。。构建树的时候数下一个就好 //疯狂超时
	题解: 思路是一样的 不过题解判断了一个cost数组，表示最小转换次数 原理在于可以有最不耗时的方法到达这个点 没必要绕着去，即转化几次
*/

func isNextWord(a string, b string) bool {

	//已经假设所有单词一样长了
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[i + 1:] == b[i + 1:]
		}
	}

	return false
}

//用cost数组替代
func contain(a string, arr *[]string) bool {
	//剪枝用，看看是否包含
	for _, s := range *arr {
		if a == s {
			return true
		}
	}

	return false
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	var ans [][]string

	//构造初始树
	wordmap := make(map[string][]string, len(wordList) + 1)
	cost := make(map[string]int, len(wordList))
	var beginnexts []string
	for i := 0; i < len(wordList); i++ {

		cost[wordList[i]] = len(wordList)

		if isNextWord(beginWord, wordList[i]) {
			beginnexts = append(beginnexts, wordList[i])
		}

		for j := i + 1; j < len(wordList); j++ {
			if isNextWord(wordList[i], wordList[j]) {
				if _, ok := wordmap[wordList[i]]; ok {
					wordmap[wordList[i]] = append(wordmap[wordList[i]], wordList[j])
				} else {
					wordmap[wordList[i]] = []string{wordList[j]}
				}

				if _, ok := wordmap[wordList[j]]; ok {
					wordmap[wordList[j]] = append(wordmap[wordList[j]], wordList[i])
				} else {
					wordmap[wordList[j]] = []string{wordList[i]}
				}
			}
		}
	}

	if _, ok := wordmap[beginWord]; !ok {
		//begin不在map里 加一下
		wordmap[beginWord] = beginnexts
	}

	minanslen := len(wordList) + 1
	//广度求解
	queue := [][]string{[]string{beginWord}}
	for len(queue) > 0 {
		curarr := queue[0]
		if len(curarr) >= minanslen {
			//已经不是最短的了
			break
		}
		queue = queue[1:]

		end := curarr[len(curarr) - 1]
		nextwordlist := wordmap[end]

		for i := 0; i < len(nextwordlist); i++ {
			next := nextwordlist[i]
			//不检查 不怕回路，回路肯定有问题
			//if contain(next, &curarr) {
			//	continue
			//}

			//使用cost数组替换回路判断
			if cost[next] > len(curarr) {
				continue
			}
			cost[next] = len(curarr)

			cu := make([]string, len(curarr))
			copy(cu, curarr)
			cu = append(cu, next)

			if next == endWord {
				ans = append(ans, cu)
				//找到第一个答案 可以更换一下剪枝用的len了
				if len(cu) < minanslen {
					minanslen = len(cu)
				}
				continue
			}

			if len(cu) < minanslen {
				queue = append(queue, cu)
			}
		}
	}

	return ans
}