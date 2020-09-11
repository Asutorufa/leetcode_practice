package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。
比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:
输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]



说明：
    用返回一个整数列表来代替打印
    n 为正整数
*/

/*
   大数问题, 防止溢出, 用字符串代替数字计算
   解决方法 -> 字符串拼接

   比如:
   输入 3


   过程(需要去掉开头的0和前缀0):
   准备 0 ... 9 单独的字符, 然后用他们组合成一个数字
   000 001 002 003 004 ... 009
   010 011 012 013 014 ... 019
              ...
   090 091 092 093 094 ... 099
   100 101 102 103 104 ... 109
              ...
   990 991 992 993 994 ... 999
*/
func printNumbers(n int) (res []int) {
	num := make([]byte, n)
	var stringBuild = &strings.Builder{}
	var nine = 0
	var start = n - 1
	var loop = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	dfs(0, n, num, stringBuild, &nine, &start, loop)
	// fmt.Println(stringBuild.String()[:stringBuild.Len()-1])
	for _, x := range strings.Split(stringBuild.String()[:stringBuild.Len()-1], ",")[1:] {
		s, _ := strconv.Atoi(x)
		res = append(res, s)
	}
	return
}

func dfs(i, n int, num []byte, stringBuild *strings.Builder, nine *int, start *int, loop []byte) {
	if i == n {
		stringBuild.WriteString(string(num[*start:]))
		stringBuild.WriteString(",")
		if n-*start == *nine {
			*start--
		}
		return
	}

	for index := range loop {
		if loop[index] == '9' {
			*nine++
		}
		num[i] = loop[index]
		dfs(i+1, n, num, stringBuild, nine, start, loop)
	}
	*nine--
}

/*
   未考虑大数问题
*/
func printNumbers2(n int) (res []int) {
	for y := float64(1); y < math.Pow(10, float64(n)); y++ {
		res = append(res, int(y))
	}
	return
}

func main() {
	fmt.Println(printNumbers(1))
	fmt.Println(printNumbers(3))

	fmt.Println(printNumbers2(1))
	fmt.Println(printNumbers2(3))
}
