package main

import "fmt"

/**
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

示例 1：
输入: n = 3
输出: 6

示例 2：
输入: n = 9
输出: 45

限制：
    1 <= n <= 10000
*/

// 逻辑符短路 这个虽然知道 但是确实想不到
func sumNums(n int) int {
	res := 0
	var f func(int) int
	f = func(n int) int {
		_ = n > 1 && f(n-1) > 0 // <- 这里
		res += n
		return res
	}
	return f(n)
}

func main() {
	fmt.Println(sumNums(3))
	fmt.Println(sumNums(9))
}
