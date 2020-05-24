package question4
/*
	4. Median of Two Sorted Arrays
	There are two sorted arrays nums1 and nums2 of size m and n respectively.
	Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
	You may assume nums1 and nums2 cannot be both empty.

	4. 寻找两个正序数组的中位数
	给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
	请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
	你可以假设 nums1 和 nums2 不会同时为空。
*/

/*
	思路: 没啥好思路，用双指针找了个中位数位置，时间复杂度 O(m + n)，不满足要求
	题解: 见到log用二分，题目转化为找第k小的数，和上面思路一致，不过这个是一片一片排除的
*/

//尝试写一下二分的方法

func find2Min(nums1 []int, nums2 []int) (int, int) {

	min1 := int(^uint(0) >> 1)
	min2 := int(^uint(0) >> 1)
	for _, i := range nums1 {
		if i < min1 {
			min2 = min1
			min1 = i
		} else if i < min2 {
			min2 = i
			break
		}
	}

	for _, i := range nums2 {
		if i < min1 {
			min2 = min1
			min1 = i
		} else if i < min2 {
			min2 = i
			break
		}
	}

	return min1, min2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	alllen := len(nums1) + len(nums2)
	var k int
	if alllen % 2 == 0 {
		k = alllen / 2 //第k个元素就是中位数，奇数情况取第k个元素，偶数情况取k和k+1
	} else {
		k = alllen / 2 + 1
	}
	//为了满足logm+n，采用二分法来找k
	//两个数组都取前k的一半，更小的那个部分可以全部认为不属于前k个元素
	for ; k > 1; {
		if len(nums1) == 0 {
			startloc := k - 1
			nums2 = nums2[startloc :]
			break
		} else if len(nums2) == 0 {
			startloc := k - 1
			nums1 = nums1[startloc :]
			break
		}

		startloc := k / 2 - 1
		var num1data int
		var num2data int
		if len(nums1) <= startloc + 1 {
			num1data = nums1[len(nums1) - 1]
		} else {
			num1data = nums1[startloc]
		}
		if len(nums2) <= startloc + 1 {
			num2data = nums2[len(nums2) - 1]
		} else {
			num2data = nums2[startloc]
		}

		if num1data <= num2data{
			if len(nums1) <= startloc + 1 {
				k -= len(nums1)
				nums1 = []int{}
				continue
			}
			nums1 = nums1[startloc + 1:]
		} else {
			if len(nums2) <= startloc + 1 {
				k -= len(nums2)
				nums2 = []int{}
				continue
			}
			nums2 = nums2[startloc + 1:]
		}

		k = k - k / 2
	}

	min1, min2 := find2Min(nums1, nums2)
	if alllen % 2 == 0 {
		return float64(min1 + min2) / 2
	} else {
		return float64(min1)
	}
}

//这个地方也能改进，与其奇偶分开处理，不如都取两个值，奇数情况拿1个，偶数情况全都要
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//
//	alllen := len(nums1) + len(nums2)
//	midloc := alllen / 2
//	iseven := false
//	if alllen % 2 == 0 {
//		iseven = true
//		midloc -= 1
//	}
//
//	nums1start := 0
//	nums2start := 0
//	var nums1now int
//	var nums2now int
//	if len(nums1) <= 0 {
//		nums1now = nums2[len(nums2) - 1]
//	} else {
//		nums1now = nums1[nums1start]
//	}
//	if len(nums2) <= 0 {
//		nums2now = nums1[len(nums1) - 1]
//	} else {
//		nums2now = nums2[nums2start]
//	}
//	addflag := false
//	midnumsum := 0
//
//	for i := 0; i < alllen; i++ {
//		if nums2start >= len(nums2) || (nums1now <= nums2now && nums1start < len(nums1)) {
//			if i == midloc || i == midloc + 1 {
//				midnumsum += nums1[nums1start]
//				if !iseven {
//					break
//				} else {
//					if !addflag {
//						addflag = true
//					} else {
//						break
//					}
//				}
//			}
//			nums1start++
//			if nums1start < len(nums1) {
//				nums1now = nums1[nums1start]
//			}
//		} else if nums1start >= len(nums1) || (nums2now < nums1now && nums2start < len(nums2)) {
//			if i == midloc || i == midloc + 1 {
//				midnumsum += nums2[nums2start]
//				if !iseven {
//					break
//				} else {
//					if !addflag {
//						addflag = true
//					} else {
//						break
//					}
//				}
//			}
//			nums2start++
//			if nums2start < len(nums2) {
//				nums2now = nums2[nums2start]
//			}
//		} else {
//			break
//		}
//	}
//
//	if iseven {
//		return float64(midnumsum) / 2
//	} else {
//		return float64(midnumsum)
//	}
//}