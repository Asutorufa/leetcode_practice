package main

import (
	"fmt"
)

/**
请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。
例如，把 9 表示成二进制是 1001，有 2 位是 1。
因此，如果输入 9，则该函数输出 2。

示例 1：

输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

示例 2：

输入：00000000000000000000000010000000
输出：1
解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。

示例 3：

输入：11111111111111111111111111111101
输出：31
解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。
*/

// 投机取巧的方法
func hammingWeight(num uint32) (res int) {
	for _, x := range fmt.Sprintf("%b", num) {
		if x == '1' {
			res++
		}
	}
	return
}

func hammingWeight2(num uint32) (res int) {
	for num != 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return
}

// 循环消去n最右边的1 <- 相当巧妙的解法(想不到)
func hammingWeight3(num uint32) (res int) {
	for num != 0 {
		res++
		num &= num - 1
	}
	return
}

func main() {
	fmt.Println(hammingWeight(9))                                  // 2
	fmt.Println(hammingWeight(0b00000000000000000000000000001011)) // 3
	fmt.Println(hammingWeight(0b00000000000000000000000010000000)) // 1
	fmt.Println(hammingWeight(0b11111111111111111111111111111101)) // 31
	fmt.Println(hammingWeight2(9))
	fmt.Println(hammingWeight2(0b00000000000000000000000000001011))
	fmt.Println(hammingWeight2(0b00000000000000000000000010000000))
	fmt.Println(hammingWeight2(0b11111111111111111111111111111101))
	fmt.Println(hammingWeight3(9))
	fmt.Println(hammingWeight3(0b00000000000000000000000000001011))
	fmt.Println(hammingWeight3(0b00000000000000000000000010000000))
	fmt.Println(hammingWeight3(0b11111111111111111111111111111101))
}
