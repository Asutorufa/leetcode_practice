package main

import (
	"fmt"
	"strings"
)

/**
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
为简单起见，标点符号和普通字母一样处理。
例如输入字符串"I am a student. "，则输出"student. a am I"。

示例 1：
输入: "the sky is blue"
输出: "blue is sky the"

示例 2：
输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

示例 3：
输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

说明：
    无空格字符构成一个单词。
    输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
    如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
*/

func reverseWords(s string) string {
	offset := 0
	l := len(s)
	arr := []string{}
	var tmp string
	for offset < l {
		tmp, offset = getWord(s, l, offset)
		if tmp == "" {
			continue
		}
		arr = append([]string{tmp}, arr...)
	}
	return strings.Join(arr, " ")
}

func getWord(s string, l, n int) (res string, offset int) {
	i := n
	j := n
	for ; i < l; i++ {
		if i > j && s[i] != ' ' {
			break
		}
		if s[i] == ' ' {
			continue
		}
		j++
	}
	return s[n:j], i
}

// 投机取巧
func reverseWords2(s string) string {
	a := strings.Split(s, " ")
	arr := []string{}
	for index := range a {
		if a[index] == "" {
			continue
		}
		arr = append([]string{a[index]}, arr...)
	}
	return strings.Join(arr, " ")
}

// 投机取巧
func reverseWords3(s string) string {
	str := strings.Fields(s)
	res := strings.Builder{}
	l := len(str) - 1
	for i := l; i >= 0; i-- {
		if i != l {
			res.WriteByte(' ')
		}
		res.WriteString(strings.TrimSpace(str[i]))
	}
	return res.String()
}

func main() {
	fmt.Println("-" + reverseWords("s string") + "-")
	fmt.Println("-" + reverseWords(" s   string ") + "-")
	fmt.Println("-" + reverseWords(" I am  a   shit. ") + "-")
	fmt.Println("-" + reverseWords("  hello world!  ") + "-")
	fmt.Println("-" + reverseWords("  ") + "-")
	fmt.Println("-" + reverseWords("") + "-")
	fmt.Println("\n Test2")
	fmt.Println("-" + reverseWords2("s string") + "-")
	fmt.Println("-" + reverseWords2(" s   string ") + "-")
	fmt.Println("-" + reverseWords2(" I am  a   shit. ") + "-")
	fmt.Println("-" + reverseWords2("  hello world!  ") + "-")
	fmt.Println("-" + reverseWords2("  ") + "-")
	fmt.Println("-" + reverseWords2("") + "-")
	fmt.Println("\n Test3")
	fmt.Println("-" + reverseWords3("s string") + "-")
	fmt.Println("-" + reverseWords3(" s   string ") + "-")
	fmt.Println("-" + reverseWords3(" I am  a   shit. ") + "-")
	fmt.Println("-" + reverseWords3("  hello world!  ") + "-")
	fmt.Println("-" + reverseWords3("  ") + "-")
	fmt.Println("-" + reverseWords3("") + "-")
}
