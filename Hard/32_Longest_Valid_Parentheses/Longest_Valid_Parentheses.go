package question32

/*
	32. 最长有效括号
	给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
*/

/*
	思路: 本来想找"()"再往外扩，后来想想还是栈方便，把匹配的干掉，栈里面剩下的就是分割用的
	题解:
		1. DP
		2. 不用结构体，直接记录一个位置信息在栈里

*/

type Node struct {
	chr rune
	loc int
}

func longestValidParentheses(s string) int {

	var stack []*Node

	for i, n := range s {

		stacklen := len(stack)

		if stacklen > 0 && n == ')' && stack[stacklen - 1].chr == '(' {
			stack = stack[:stacklen - 1]
			continue
		}

		stack = append(stack, &Node{n, i})
	}

	maxlen := 0
	start := 0

	for _, n := range stack {
		tlen := n.loc - start
		if tlen > maxlen {
			maxlen = tlen
		}
		start = n.loc + 1
	}

	// 到结尾
	tlen := len(s) - start
	if tlen > maxlen {
		maxlen = tlen
	}

	return maxlen
}