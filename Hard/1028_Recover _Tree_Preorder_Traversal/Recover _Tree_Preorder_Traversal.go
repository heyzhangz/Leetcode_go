package question1028

/*
	1028. 从先序遍历还原二叉树
	我们从二叉树的根节点 root 开始进行深度优先搜索。
	在遍历中的每个节点处，我们输出 D 条短划线（其中 D 是该节点的深度），然后输出该节点的值。（如果节点的深度为 D，则其直接子节点的深度为 D + 1。根节点的深度为 0）。
	如果节点只有一个子节点，那么保证该子节点为左子节点。
	给出遍历输出 S，还原树并返回其根节点 root。
*/

/*
	思路: 这是个递归 出错的原因是忘了两遍 左右子树都要加一次
	题解:
		用栈改迭代
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func addChild(root *TreeNode, leaf *TreeNode) {

	if root.Left == nil {
		root.Left = leaf
	} else {
		root.Right = leaf
	}
}

func buildTree(S string, pos int, predepth int, root *TreeNode) int {

	if len(S) <= pos {
		return pos
	}

	curval, curpos, curdepth := 0, pos, 0
	breakflag := false
	for ; curpos < len(S); curpos++ {
		if S[curpos] == '-' {
			if breakflag {
				break
			}
			curdepth++
		} else {
			curval *= 10
			curval += int(S[curpos]) - int('0')
			breakflag = true
		}
	}

	if curdepth > predepth {
		curnode := TreeNode{Val: curval, Left: nil, Right: nil}
		addChild(root, &curnode)
		pos = buildTree(S, curpos, curdepth, &curnode)
		if pos != curpos {
			pos = buildTree(S, pos, curdepth, &curnode)
		}
	}

	return pos
}

func recoverFromPreorder(S string) *TreeNode {

	if len(S) <= 0 {
		return nil
	}

	rootnum, pos := 0, 0
	for _, n := range S {
		if n == '-' {
			break
		}

		rootnum *= 10
		rootnum += int(n) - int('0')
		pos++
	}

	root := TreeNode{Val: rootnum, Left: nil, Right: nil}
	pos = buildTree(S, pos, 0, &root)
	pos = buildTree(S, pos, 0, &root)

	return &root
}