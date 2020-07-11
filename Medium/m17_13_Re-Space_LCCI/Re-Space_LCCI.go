package questionm02_05

import (
    "math"
    "strings"
)

/*
	面试题 17.13. 恢复空格
	哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。像句子"I reset the computer. It still didn’t boot!"已经变成了"iresetthecomputeritstilldidntboot"。在处理标点符号和大小写之前，你得先把它断成词语。当然了，你有一本厚厚的词典dictionary，不过，有些词没在词典里。假设文章用sentence表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。

    注意：本题相对原题稍作改动，只需返回未识别的字符数
 */

/*
	思路: 没啥思路；看题解用到了字典树，那么剪枝+深度应该ok，不过有重复判断应该..递归 + hash
	题解:
        这个从后往前的dp有点东西
*/

type TrieTree struct {
    dic map[rune]*TrieTree
    end bool
}

func (root *TrieTree)addWord(word string) {

    call := root
    word2 := []rune(word)
    for i := len(word2) - 1; i >= 0; i-- {
        n := word2[i]
        if call.dic == nil {
            call.dic = make(map[rune]*TrieTree)
            call.dic[n] = &TrieTree{dic: nil, end: false}
        } else {
            if _, ok := call.dic[n]; !ok {
                call.dic[n] = &TrieTree{dic: nil, end: false}
            }
        }
        call = call.dic[n]
    }
    call.end = true
}

func (root *TrieTree)next(letter rune) bool {

    _, ok := root.dic[letter]
    return ok
}

func (root *TrieTree)nextEnd(letter rune) bool {

    next := root.dic[letter]

    return next.end
}

func respace(dictionary []string, sentence string) int {

    dictree := TrieTree{nil, false}
    for _, word := range dictionary {
        dictree.addWord(word)
    }

    dp := make([]int, len(sentence) + 1)
    for i := 1; i < len(dp); i++ {
        dp[i] = math.MaxInt32
    }

    sentence2 := []rune(sentence)
    for i := 1; i <= len(sentence2); i++ {

        dp[i] = dp[i - 1] + 1

        curpos := &dictree
        for j := i; j >= 1; j-- {
            t := sentence2[j - 1]
            if !curpos.next(t) {
                break
            } else if curpos.nextEnd(t) {
                dp[i] = min(dp[i], dp[j - 1])
            }

            if dp[i] == 0 {
                break
            }
            curpos = curpos.dic[t]
        }
    }

    return dp[len(dp) - 1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func respace_onlydp(dictionary []string, sentence string) int {

    dp := make([]int, len(sentence) + 1)
    dp[len(sentence)] = 0

    for i := len(sentence) - 1; i >= 0; i-- {
        b := sentence[i:]
        min := dp[i + 1] + 1

        for j := 0; j < len(dictionary); j++ {

            if strings.HasPrefix(b, dictionary[j]) {

                if min > dp[i + len(dictionary[j])] {
                    min = dp[i + len(dictionary[j])]
                }
            }
        }

        dp[i] = min
    }
    return dp[0]
}