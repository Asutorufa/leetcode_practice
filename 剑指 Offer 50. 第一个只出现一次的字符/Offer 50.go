package main

import (
	"bytes"
	"fmt"
)

/**
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例:
s = "abaccdeff"
返回 "b"

s = ""
返回 " "

限制：
0 <= s 的长度 <= 50000
*/

// 简单题... 没想到会这么简单...
func firstUniqChar(s string) byte {
	hash := make(map[byte]int)
	for index := range s {
		hash[s[index]]++
	}

	for index := range s {
		if hash[s[index]] > 1 {
			continue
		}
		return s[index]
	}
	return ' '
}

func firstUniqChar2(s string) byte {
	hash := make(map[byte]int)
	arr := make([]byte, 0)
	for index := range s {
		if bytes.IndexByte(arr, s[index]) == -1 {
			arr = append(arr, s[index])
		}
		hash[s[index]]++
	}

	for index := range arr {
		if hash[arr[index]] > 1 {
			continue
		}
		return arr[index]
	}
	return ' '
}

func main() {
	fmt.Println(string(firstUniqChar("abaccdeff")))
	fmt.Println(string(firstUniqChar2("abaccdeff")))
}
