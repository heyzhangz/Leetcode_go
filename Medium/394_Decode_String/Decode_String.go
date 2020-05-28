package question394

import (
	"regexp"
	"strconv"
)

/*
	394. Decode String
	Given an encoded string, return its decoded string.
	The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.
	You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.
	Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].

	394. 字符串解码
	给定一个经过编码的字符串，返回它解码后的字符串。
	编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
	你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
	此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 */

/*
	思路: 这个跟加减乘除那个没区别，就是个栈。。这里考虑用正则拆开字符串的话其实不用考虑"["
	题解:
*/

func decodeString(s string) string {

	re := regexp.MustCompile(`(?:\d+)|(?:[a-zA-Z]+)|(?:\[)|(?:])`)
	matched := re.FindAllString(s, -1)

	var stack []string

	for _, s := range matched {

		if s == "]" {
			stacklen := len(stack)
			curstr := ""
			i := stacklen - 1
			for ; i > 0; i-- {
				if stack[i] == "[" {
					break
				}
				curstr = stack[i] + curstr
			}

			tempstr := ""
			num, _ := strconv.Atoi(stack[i - 1])
			stack = stack[: i - 1]
			for i := 0; i < num; i++ {
				tempstr += curstr
			}
			stack = append(stack, tempstr)
		} else {
			stack = append(stack, s)
		}
	}

	ans := ""
	if stack != nil {
		for _, s := range stack {
			ans += s
		}
	}

	return ans
}