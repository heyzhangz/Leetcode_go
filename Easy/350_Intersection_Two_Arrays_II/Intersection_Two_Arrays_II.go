package question350

import "math"

/*
	350. 两个数组的交集 II
	给定两个数组，编写一个函数来计算它们的交集。

*/

/*
	思路: hash可以，想了想bool数组更省点空间
	题解:
		hash和排序后双指针
*/

func intersect(nums1 []int, nums2 []int) []int {
    
    var less, more *[]int
    if len(nums1) > len(nums2) {
        more = &nums1
        less = &nums2
    } else {
        more = &nums2
        less = &nums1
    }

    id := make([]bool, len(*less))
    for _, n := range *more {
        for j, m := range *less {
            if n == m && !id[j] {
                id[j] = true
                break
            }
        }
    }

    var ans []int
    for i, n := range id {
        if n {
            ans = append(ans, (*less)[i])
        }
    }

    return ans
}