package question315

type BSTnode struct {
	value int
	leftcount int // 比它小的数
	left *BSTnode
	right *BSTnode // 左小右大
}

//值传递还是慢 用指针
func (root *BSTnode) addNode(node *BSTnode) int {

	lessnum := 0

	if node.value > root.value {
		lessnum += root.leftcount + 1
		if root.right == nil {
			root.right = node
		} else {
			lessnum += root.right.addNode(node)
		}
	} else {
		root.leftcount += 1
		if root.left == nil {
			root.left = node
		} else {
			lessnum += root.left.addNode(node)
		}
	}

	return lessnum
}

func countSmaller_BST(nums []int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	ans := make([]int, len(nums))
	ans[len(nums) - 1] = 0
	root := BSTnode{nums[len(nums) - 1], 0, nil, nil}

	for i := len(nums) - 2; i >= 0; i-- {

		node := BSTnode{nums[i], 0, nil, nil}
		lessnum := root.addNode(&node)
		ans[i] = lessnum
	}

	return ans
}