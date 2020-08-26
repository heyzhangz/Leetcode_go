package question337

/*
	17. 电话号码的字母组合
	给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
	给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 */

/*
	思路: 回溯一哈
	题解:
		same
*/

var phoneMap = map[string]string {
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	ans := []string{}

	backtrack(&ans, digits, 0, "")
	return ans
}

func backtrack(ans *[]string, digits string, index int, combination string) {

	if index == len(digits) {
		*ans = append(*ans, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]

		for _, i := range letters {
			backtrack(ans, digits, index + 1, combination + string(i))
		}
	}

}