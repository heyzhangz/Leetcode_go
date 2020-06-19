package question125

/*
	125. 验证回文串
	给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
	说明：本题中，我们将空字符串定义为有效的回文串。

*/

/*
	思路: 双指针
	题解: 

*/

func isValid(a byte) bool {

    if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
        return true
    }

    return false
}

func isPalindrome(s string) bool {

    if len(s) <= 1 {
        return true
    }

    for i, j := 0, (len(s) - 1); i < j && i < len(s) && j >= 0; {

        for !isValid(s[i]) && i < j {
            i++
        }
        for !isValid(s[j]) && i < j {
            j--
        }

        if i >= j || i >= len(s) || j < 0 {
            break
        }

        var a, b byte
        a = s[i]
        b = s[j]
        if a >= 'A' && a <= 'Z' {
            a = a - 'A' + 'a'
        }

        if b >= 'A' && b <= 'Z' {
            b = b - 'A' + 'a'
        }

        if a != b {
            return false
        }

        i++
        j--
    }

    return true
}