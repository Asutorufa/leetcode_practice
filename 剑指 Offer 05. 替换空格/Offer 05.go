package main

import (
	"log"
	"strings"
)

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."



限制：

0 <= s 的长度 <= 10000
*/

func replaceSpace3(s string) string {
	subStr := 0
	for index := range s {
		if s[index] == ' ' {
			subStr++
		}
	}
	str := make([]byte, len(s)+(len("%20")-len(" "))*subStr)
	nowSize := 0
	for index := range s {
		if s[index] == ' ' {
			str[nowSize] = '%'
			str[nowSize+1] = '2'
			str[nowSize+2] = '0'
			nowSize += 3
		} else {
			str[nowSize] = s[index]
			nowSize += 1
		}
	}
	return string(str)
}

func replaceSpace(s string) string {
	str := ""
	for index := range s {
		if s[index] == ' ' {
			str = str + "%20"
		} else {
			str = str + string(s[index])
		}
	}
	return str
}

// 暴力 但是这样做就没有意义了
func replaceSpace2(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
func main() {
	str := "We are happy."
	log.Println(replaceSpace(str))
}
