package question15

import "sort"

/*
	15. 三数之和
	给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
	注意：答案中不可以包含重复的三元组。
 */

/*
	思路: 选定一个数，那么在剩余数组种这就是一个二数之和的问题。。注意去重 去重代码太臭了。。
	题解:
		题解通过排序来去重 采用双指针 也是O(n2)
*/

func twoSum(nums []int, sum int) [][]int {

	statusmap := make(map[int]bool, len(nums))

	for _, n := range nums {

		rest := sum - n
		if _, ok := statusmap[n]; ok {
			//找到差值
			statusmap[n] = true
		} else {
			if _, ok := statusmap[rest]; !ok {
				statusmap[rest] = false
			}
		}
	}

	var ans [][]int
	for k, v := range statusmap {
		if v {
			ans = append(ans, []int{k, sum - k})
		}
	}

	return ans
}

func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return [][]int{}
	}

	var ans [][]int
	//去重 只要两个值定了 第三个值就定了
	temp := make(map[int]map[int]bool)

	for i, n := range nums[0:len(nums) - 2] {
		res := twoSum(nums[i + 1:], 0 - n)
		if len(res) > 0 {
			for _, nn := range res {

				if _, ok := temp[n][nn[0]]; ok {
					continue
				}

				if temp[n] == nil {
					temp[n] = map[int]bool {
						nn[0]: true,
						nn[1]: true,
					}
				} else {
					temp[n][nn[0]] = true
					temp[n][nn[1]] = true
				}

				if temp[nn[0]] == nil {
					temp[nn[0]] = map[int]bool {
						n: true,
						nn[1]: true,
					}
				} else {
					temp[nn[0]][n] = true
					temp[nn[0]][nn[1]] = true
				}

				if temp[nn[1]] == nil {
					temp[nn[1]] = map[int]bool {
						n: true,
						nn[0]: true,
					}
				} else {
					temp[nn[1]][n] = true
					temp[nn[1]][nn[0]] = true
				}

				nn = append(nn, n)
				ans = append(ans, nn)
			}
		}
	}

	return ans
}

func threeSum_ans(nums []int) [][]int {

	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}