package main

import (
	"fmt"
)

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：

输入：n = 2
输出：2

示例 2：

输入：n = 7
输出：21

示例 3：

输入：n = 0
输出：1

提示：

    0 <= n <= 100
*/

func numWays(n int) (b int) { // Dynamic programming => less time, less memory
	a := 1
	b = 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b // a = last b, b = last a + last b
		b %= 1000000007
	}
	return b
}

/*
   设跳上 n 级台阶有 f(n)种跳法。在所有跳法中，青蛙的最后一步只有两种情况： 跳上 1 级或 2 级台阶。
       当为 1 级台阶： 剩 n−1 个台阶，此情况共有 f(n−1) 种跳法；
       当为 2 级台阶： 剩 n−2 个台阶，此情况共有 f(n−2) 种跳法。
   f(n) 为以上两种情况之和，即 f(n)=f(n−1)+f(n−2) ，以上递推性质为斐波那契数列。本题可转化为 求斐波那契数列第 n 项的值 ，与 面试题10- I. 斐波那契数列 等价，唯一的不同在于起始数字不同。
       青蛙跳台阶问题： f(0)=1 , f(1)=1 , f(2)=2；
       斐波那契数列问题： f(0)=0 , f(1)=1 , f(2)=1 。
*/

func main() {
	fmt.Println(numWays(0))
	fmt.Println(numWays(2))
	fmt.Println(numWays(7))
}
