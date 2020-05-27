package question974

/*
	974. Subarray Sums Divisible by K
	Given an array A of integers, return the number of (contiguous, non-empty) subarrays that have a sum divisible by K.

	974. 和可被 K 整除的子数组
	给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。
 */

/*
	思路: 第一反应动规 D[i, i + len] 的状态可以利用 D[i, i + len - 1] =。=其实就是个遍历，还不如不存状态直接暴力判断
	题解: 通常，涉及连续子数组问题的时候，我们使用前缀和来解决。 O(n）
		 P[i, j] = P[j] - P[i] 那就保证P[j] 和 P[i] 的余数相同就好
*/

func subarraysDivByK(A []int, K int) int {

	prefix := make(map[int]int) // map[余数] loc
	prefix[0] = 1
	ans := 0

	tempsum := 0
	for _, n := range A {
		tempsum += n
		remainder := tempsum % K
		if remainder < 0 {
			remainder = remainder + K
		}
		if _, ok := prefix[remainder]; ok {
			ans += prefix[remainder]
			prefix[remainder] += 1
		} else {
			prefix[remainder] = 1
		}
	}

	return ans
}

//func subarraysDivByK(A []int, K int) int {
//
//	seqstatus := make([][]int, len(A))
//	for i, _ := range seqstatus {
//		seqstatus[i] = make([]int, len(A))
//	}
//
//	ans := 0
//
//	for i, n := range A {
//		seqstatus[i][i] = n
//		if n % K == 0 {
//			ans++
//		}
//	}
//
//	for length := 1; length < len(A); length++ {
//
//		for i := 0; i < len(A); i++ {
//			if i + length >= len(A) {
//				break
//			}
//			seqstatus[i][i + length] = seqstatus[i][i + length - 1] + seqstatus[i + length][i + length]
//			if seqstatus[i][i + length] % K == 0 {
//				ans++
//			}
//		}
//	}
//
//	return ans
//}