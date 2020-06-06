package question128

/*
	并查集 核心思路是 union 就是将相邻节点合起来
*/

func longestConsecutive_UF(nums []int) int {

	//因为是连续的 所以这里集合就不用数组表示了 用 start end，把相邻节点合起来
	unionfind := make(map[int]int, len(nums))

	for _, n := range nums {
		unionfind[n] = n //初始化 指向自己
	}


	maxlen := 0
	for _, n := range nums {

		if _, ok := unionfind[n]; !ok {
			continue
		}

		curun := n
		for true {
			if _, ok := unionfind[curun + 1]; ok {
				//uniob操作，和下一个
				unionfind[n] = unionfind[curun + 1]
				delete(unionfind, curun + 1)
				curun = unionfind[n]
			} else {
				curlen := unionfind[n] - n + 1
				if curlen > maxlen {
					maxlen = curlen
				}
				break
			}
		}
	}

	return maxlen
}