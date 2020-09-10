package main

import "fmt"

/**
实现函数double Power(double base, int exponent)，求base的exponent次方。
不得使用库函数，同时不需要考虑大数问题。

示例 1:
输入: 2.00000, 10
输出: 1024.00000

示例 2:
输入: 2.10000, 3
输出: 9.26100

示例 3:
输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25

说明:
    -100.0 < x < 100.0
    n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/

/**
* quick pow
* 3^10 -> 3^(0b1010) -> 3^(0b0)*3^(0b10)*3^(0b0)*3^(0b1000) = 3^(0b1000)*3^(0b0)*3^(0b10)*3^(0b0)
 */
func myPow(x float64, n int) (res float64) {
	if x == 0 {
		return 0
	}
	res = 1
	if n < 0 {
		x = 1 / x
		n = -n
	}
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		// fmt.Println(x, n&1, res)
		x *= x
		n >>= 1
	}
	return
}

//  Brute Force <- 超出时间限制
func myPow2(x float64, n int) (res float64) {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	res = x
	s := n >= 0
	if !s {
		n *= -1
	}
	for y := 0; y < n-1; y++ {
		res *= x
	}
	if !s {
		return 1 / res
	}
	return
}

func main() {
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2.0, -2))
	fmt.Println(myPow(3.0, 10))
	fmt.Println(myPow2(2.0, 10))
	fmt.Println(myPow2(2.1, 3))
	fmt.Println(myPow2(2.0, -2))
}
