package questionm16

import "strings"

/*
	面试题 16.18. 模式匹配
	你有两个字符串，即pattern和value。 pattern字符串由字母"a"和"b"组成，用于描述字符串中的模式。
	例如，字符串"catcatgocatgo"匹配模式"aabab"（其中"cat"是"a"，"go"是"b"），该字符串也匹配像"a"、"ab"和"b"这样的模式。但需注意"a"和"b"不能同时表示相同的字符串。编写一个方法判断value字符串是否匹配pattern字符串。
 */

/*
	思路: a b匹配模式的长度其实是固定的 总长要和字符串相等 好恶心！这犄角旮旯的case太多了
	题解:
		也是长度，这题考的边界条件
*/

func patternMatching(pattern string, value string) bool {

	acount := 0
	bcount := 0
	startcount := 0
	conflag := true
	for _, n := range pattern {
		if n == 'a' {
			acount++
		} else {
			bcount++
		}

		if conflag {
			if rune(pattern[0]) == n {
				startcount++
			} else {
				conflag = false
			}
		}
	}

	if value == "" {
		if acount != 0 && bcount != 0 {
			return false
		}
		return true
	}

	if acount == 0 && bcount == 0 {
		if value == "" {
			return true
		}
		return false
	}

	if acount == 0 {
		pattern = strings.ReplaceAll(pattern, "b", "a")
		acount = bcount
		bcount = 0

	}

	for alen := 0; alen <= len(value); alen++ {

		asumlen := alen * acount
		bsumlen := len(value) - asumlen
		blen := 0

		if bsumlen < 0 {
			continue
		}

		if bcount != 0 {
			if bsumlen % bcount != 0 {
				continue
			}

			blen = bsumlen / bcount
		} else {
			if bsumlen != 0 {
				continue
			}
		}

		astr, bstr := "", ""
		if pattern[0] == 'a' {
			start := alen*startcount
			astr = value[:alen]
			bstr = value[start:start+blen]
		} else {
			start := blen*startcount
			bstr = value[:blen]
			astr = value[start:start+alen]
		}


		pos := 0
		stopflag := false
		for _, n := range pattern {
			if n == 'a' {
				if value[pos:pos+alen] != astr {
					stopflag = true
					break
				}
				pos += alen
			} else {
				if value[pos:pos+blen] != bstr {
					stopflag = true
					break
				}
				pos += blen
			}
		}

		if stopflag || pos != len(value) {
			continue
		}

		return true
	}

	return false
}