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



type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func isSubSymmetric(left *TreeNode, right *TreeNode) bool {

	if left == nil {
		if right != nil {
			return false
		} else {
			return true
		}
	}

	if right == nil {
		//left已经判断过了
		return false
	}

	if right.Val != left.Val {
		return false
	}

	return isSubSymmetric(left.Right, right.Left) && isSubSymmetric(left.Left, right.Right)
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isSubSymmetric(root.Left, root.Right)
}

