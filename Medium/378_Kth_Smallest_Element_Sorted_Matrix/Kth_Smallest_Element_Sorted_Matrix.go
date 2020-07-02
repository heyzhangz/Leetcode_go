package question378

/*
	378. 有序矩阵中第K小的元素
	给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
	请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
 */

/*
	思路: 这题帮忙做了归并排序的一步
	题解: 二分？？
*/

func kthSmallest(matrix [][]int, k int) int {

	length := len(matrix)

	if length == 0 {
		return 0
	}

	//!!!! 注意数字拷贝
	arr := make([]int, length * length)
	for i, n := range matrix[0] {
		arr[i] = n
	}
	nowlen := length

	for _, line := range matrix[1 : ] {

		temparr := make([]int, nowlen)
		copy(temparr, arr[ : nowlen])
		aptr, bptr, cptr := 0, 0, 0

		for aptr < nowlen && bptr < length {
			if temparr[aptr] > line[bptr] {
				arr[cptr] = line[bptr]
				bptr++
			} else {
				arr[cptr] = temparr[aptr]
				aptr++
			}
			cptr++
		}

		for aptr < nowlen {
			arr[cptr] = temparr[aptr]
			aptr++
			cptr++
		}

		for bptr < length {
			arr[cptr] = line[bptr]
			bptr++
			cptr++
		}

		nowlen += length
	}

	return arr[k - 1]
}