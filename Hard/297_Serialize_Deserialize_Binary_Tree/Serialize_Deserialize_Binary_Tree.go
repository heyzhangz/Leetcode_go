package question297

import (
	"strconv"
	"strings"
)

/*
	297. 二叉树的序列化与反序列化
	序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
	请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
*/

/*
	思路: 这。。前序中序遍历和恢复？
	题解:
		递归层次序遍历也ok
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Codec struct {
	preorder string
	inorder string
	root *TreeNode
}

func Constructor() Codec {
	return Codec{preorder: "", inorder: "", root: nil}
}

func printPreTree(root *TreeNode, order *string, remap map[int]int) {

	if root != nil {
		if _, ok := remap[root.Val]; ok {
			remap[root.Val]++
		} else {
			remap[root.Val] = 0
		}
		tm := remap[root.Val]
		*order += strconv.Itoa(root.Val) + "_" + strconv.Itoa(tm) + ","
		printPreTree(root.Left, order, remap)
		printPreTree(root.Right, order, remap)
	}
}

func printInTree(root *TreeNode, order *string, remap map[int]int) {

	if root != nil {
		if _, ok := remap[root.Val]; ok {
			remap[root.Val]++
		} else {
			remap[root.Val] = 0
		}
		tm := remap[root.Val]
		printInTree(root.Left, order, remap)
		*order += strconv.Itoa(root.Val) + "_" + strconv.Itoa(tm) + ","
		printInTree(root.Right, order, remap)
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	printInTree(root, &this.inorder, map[int]int{})
	printPreTree(root, &this.preorder, map[int]int{})
	return this.inorder + "||" + this.preorder
}

func findLoc(val string, inorder *[]string) int {

	for i, c := range *inorder {
		if c == val {
			return i
		}
	}

	return -1
}

func buildTree(preorder []string, inorder []string) *TreeNode {

	if len(preorder) != len(inorder) {
		return nil
	}

	if len(preorder) == 0 {
		return nil
	}

	rootval := preorder[0]
	//题设没重复节点
	rootlocin := findLoc(rootval, &inorder)

	tv, _ := strconv.Atoi(strings.Split(rootval, "_")[0])
	return &TreeNode{tv,
		buildTree(preorder[1 : rootlocin + 1], inorder[0 : rootlocin]),
		buildTree(preorder[rootlocin + 1 : ], inorder[rootlocin + 1 :]),
	}
}

func changetoArray(data string) []int {
	datas := strings.Split(data, ",")
	var ans []int
	for _, n := range datas {
		if n == "" {
			continue
		}
		t, _ := strconv.Atoi(n)
		ans = append(ans, t)
	}

	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	datas := strings.Split(data, "||")
	this.inorder = datas[0]
	this.preorder = datas[1]

	inorderarr := strings.Split(this.inorder, ",")
	inorderarr = inorderarr[:len(inorderarr) - 1]
	preorderarr := strings.Split(this.preorder, ",")
	preorderarr = preorderarr[:len(preorderarr) - 1]

	this.root = buildTree(preorderarr, inorderarr)

	return this.root
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */