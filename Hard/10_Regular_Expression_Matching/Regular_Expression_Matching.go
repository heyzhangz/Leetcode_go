package question10

/*
	10. 正则表达式匹配
	给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的那一个元素
	所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
	说明:
	s 可能为空，且只包含从 a-z 的小写字母。
	p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
*/

/*
	思路: 好像也没有必要用栈 直接扫一遍 p 就好，俩指针..不对，没法处理a*b*aa
 		 构建个简单语法树
	题解: DP ..
		https://leetcode-cn.com/problems/regular-expression-matching/solution/zheng-ze-biao-da-shi-pi-pei-by-leetcode-solution/
*/

//*语法最多是个二叉树 带*可以指向自己，不带*不行，没next就ok
type Node struct {
	pos int
	value byte
	mine *Node
	next *Node
}

func (node *Node) setPos(pos int) {

	node.pos = pos
}

func (node *Node) setNext(next *Node) {

	node.next = next
}

func (node *Node) setMine() {

	node.mine = node
}

func getGraph(s string) *Node {

	root := new(Node)
	root.value = s[0]
	i := 1
	if i < len(s) && s[i] == '*' {
		root.setMine()
		i++
	}

	prenode := root
	for ; i < len(s); i++ {
		cur := s[i]
		curnode := new(Node)
		curnode.value = cur
		if i + 1 < len(s) && s[i + 1] == '*' {
			curnode.setMine()
			i++
		}
		prenode.setNext(curnode)
		prenode = curnode
	}

	return root
}

func isMatch(s string, p string) bool {

	if len(p) == 0 {
		if len(s) == 0 {
			return true
		} else {
			return false
		}
	}

	if len(s) == 0 {
		if p[len(p) - 1] != '*' {
			return false
		}
	}

	if p[len(p) - 1] != '*' && p[len(p) - 1] != '.' && p[len(p) - 1] != s[len(s) - 1] {
		return false
	}

	root := getGraph(p)
	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if node.mine != nil && node.next != nil {
			rightnode := new(Node)
			*rightnode = *(node.next)
			rightnode.setPos(node.pos)
			stack = append(stack, rightnode)
		}

		if node.pos >= len(s) {
			if node.next == nil && node.mine != nil {
				return true
			}
			continue
		}

		if node.value == '.' || s[node.pos] == node.value {

			if node.next == nil && node.pos == len(s) - 1 {
				return true
			}

			if node.mine != nil {
				leftnode := new(Node)
				*leftnode = *(node.mine)
				leftnode.setPos(node.pos + 1)
				stack = append(stack, leftnode)
			}

			if node.next != nil {
				rightnode := new(Node)
				*rightnode = *(node.next)
				rightnode.setPos(node.pos + 1)
				stack = append(stack, rightnode)
			}
		}
	}

	return false
}

func isMatch_ans(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m + 1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n + 1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j - 1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}