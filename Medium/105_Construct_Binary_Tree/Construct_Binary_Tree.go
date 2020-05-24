package question105

/*
	105. Construct Binary Tree from Preorder and Inorder Traversal
	Given preorder and inorder traversal of a tree, construct the binary tree.

	105. 从前序与中序遍历序列构造二叉树
	根据一棵树的前序遍历与中序遍历构造二叉树。
	注意:你可以假设树中没有重复的元素。
*/

/*
	思路: 递归，前序确定root，然后在中序分为两棵子树，再依次求解。。这个比较简单
	思路2: 改迭代
*/


//Tips : 前序 中序指的是啥时候变量root，=。=惊呆了这都忘了
type TreeNode struct {
     Value int
     Left *TreeNode
     Right *TreeNode
}

func findLoc(val int, inorder *[]int) int {

	for i, c := range *inorder {
		if c == val {
			return i
		}
	}

	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) != len(inorder) {
		return nil
	}

	if len(preorder) == 0 {
		return nil
	}

	rootval := preorder[0]
	//题设没重复节点
	rootlocin := findLoc(rootval, &inorder)

	return &TreeNode{rootval,
		buildTree(preorder[1 : rootlocin + 1], inorder[0 : rootlocin]),
		buildTree(preorder[rootlocin + 1 : ], inorder[rootlocin + 1 :]),
	}
}