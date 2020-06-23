package question67

import "strconv"

/*
	67. 二进制求和
	给你两个二进制字符串，返回它们的和（用二进制表示）。
	输入为 非空 字符串且只包含数字 1 和 0。

*/

/*
	思路: 偷懒用strconv果然超长了。。老老实实做
	题解:
		golang没py的高精度运算
	把 aa 和 bb 转换成整型数字 xx 和 yy，在接下来的过程中，xx 保存结果，yy 保存进位。
	当进位不为 00 时
		计算当前 xx 和 yy 的无进位相加结果：answer = x ^ y
		计算当前 xx 和 yy 的进位：carry = (x & y) << 1
		完成本次循环，更新 x = answer，y = carry

*/

func addBinary(a string, b string) string {

	ans := ""
	i := len(a) - 1
	j := len(b) - 1
	var up uint8 = 0
	for i >= 0 || j >= 0 {
		var ta uint8 = 0
		var tb uint8 = 0
		if i >= 0 {
			ta = a[i] - '0'
			i--
		}
		if j >= 0 {
			tb = b[j] - '0'
			j--
		}

		var sum uint8 = ta + tb + up
		up = sum >> 1
		sum %= 2

		ans = strconv.Itoa(int(sum)) + ans

	}

	if up > 0 {
		ans = strconv.Itoa(int(up)) + ans
	}

	return ans
}