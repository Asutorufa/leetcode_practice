package main

import (
	"fmt"
)

/*
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.

斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：

输入：n = 2
输出：1

示例 2：

输入：n = 5
输出：5

提示：

    0 <= n <= 100
*/

func fib1(n int) int { // more time, more memory
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fib(n int) int { // Dynamic programming -> less time
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := make([]int, n+1)
	a[0] = 0
	a[1] = 1
	for i := 2; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
		a[i] %= 1000000007
	}
	return a[n]
}

func fib2(n int) int { // Dynamic programming => less time, less memory
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
		b %= 1000000007
	}
	return b
}

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
	fmt.Println(fib(43))
	fmt.Println(fib(45))
	fmt.Println(fib2(2))
	fmt.Println(fib2(5))
	fmt.Println(fib2(43))
	fmt.Println(fib2(45))
}
