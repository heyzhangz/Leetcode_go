package question128

/*
	128. 最长连续序列
	给定一个未排序的整数数组，找出最长连续序列的长度。
	要求算法的时间复杂度为 O(n)。
*/

/*
	思路: 找最大然后桶排序也挺好。。。肯定要占空间了 hash表保存一下当前值和下一个值，最后找一下key
	题解:
		1. 自己思路问题在于 应该同时向前向后找，所以这里应该标识的是这个值是否被访问过 hashset
		2. 并查集 再写写
*/

func lookpre(longmap *map[int]bool, num int) int {

	curlen := 0

	for true {
		//如果有重复的 那之前已经找完了 这里一定是没找过的
		num--
		if _, ok := (*longmap)[num]; ok {
			(*longmap)[num] = true
			curlen++
		} else {
			break
		}
	}

	return curlen
}

func looknext(longmap *map[int]bool, num int) int {

	curlen := 0

	for true {
		//如果有重复的 那之前已经找完了 这里一定是没找过的
		num++
		if _, ok := (*longmap)[num]; ok {
			(*longmap)[num] = true
			curlen++
		} else {
			break
		}
	}

	return curlen
}

func longestConsecutive(nums []int) int {

	longmap := make(map[int]bool, len(nums))

	//第一遍循环 初始化map
	for _, n := range nums {
		longmap[n] = false
	}

	//开始查找
	maxlen := 0
	for k, v := range longmap {
		//已经遍历过了 跳过
		//题解此处优化 v || _, ok := longmap[n]; ok 把不是开头的过滤掉
		if v {
			continue
		}

		curlen := 1
		curlen += lookpre(&longmap, k)
		curlen += looknext(&longmap, k)

		if curlen > maxlen {
			maxlen = curlen
		}
	}

	return maxlen
}