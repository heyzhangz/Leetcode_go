package question101

/*
	101. 对称二叉树
	给定一个二叉树，检查它是否是镜像对称的。
	例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
	但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3
 */

/*
	思路: 这一看就是个递归嘛！尝试改个迭代
	题解:
*/

func isSymmetricIter(root *TreeNode) bool {

	if root == nil {
		return true
	}

	//队列
	var queue []*TreeNode
	queue = append(queue, root.Left, root.Right)

	for left, right := queue[0], queue[1]; len(queue) > 1; {
		left, right = queue[0], queue[1]
		queue = queue[2 : ]

		if left == nil {
			if right == nil {
				continue
			} else {
				return false
			}
		}

		if right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}

	return true
}

