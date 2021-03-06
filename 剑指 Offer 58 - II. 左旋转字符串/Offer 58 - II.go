package main

import (
	"fmt"
	"strings"
)

/**
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
请定义一个函数实现字符串左旋转操作的功能。
比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

示例 1：
输入: s = "abcdefg", k = 2
输出: "cdefgab"

示例 2：
输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"

限制：
    1 <= k < s.length <= 10000
*/

// 切片取巧
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

// 新建字符串
func reverseLeftWords2(s string, n int) string {
	l := len(s)
	str := strings.Builder{}
	for i := n; i < l; i++ {
		str.WriteByte(s[i])
	}
	for i := 0; i < n; i++ {
		str.WriteByte(s[i])
	}
	return str.String()
}

func main() {
	fmt.Println("Test Slice")
	fmt.Println(reverseLeftWords("abcdefg", 2))
	fmt.Println(reverseLeftWords("lrloseumgh", 6))
	fmt.Println("\nTest Range")
	fmt.Println(reverseLeftWords2("abcdefg", 2))
	fmt.Println(reverseLeftWords2("lrloseumgh", 6))
}
