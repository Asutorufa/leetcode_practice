package main

import (
	"fmt"
)

/**
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

示例:
输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。

说明:
    1 是丑数。
    n 不超过1690。
*/

// 数学规律 + dp
/**
根据题意， 一个丑数必然可以写为 A0∗A1∗.....∗A(n−1)∗An, 其中 A∈[2,3,5]。
那么这个丑数也可以写为 (A0∗A1∗.....∗A(n−1))∗An, (A0∗A1∗.....∗A(n−1)) 也是个丑数，
而 An∈[2,3,5]， 所以一个丑数乘以 2， 3， 5 之后， 一定还是一个丑数。

并且如果从丑数序列首个元素开始 *2 *3 *5 来计算， 丑数序列也是不会产生漏解的。

完整题解: https://leetcode-cn.com/problems/chou-shu-lcof/solution/chou-shu-ii-qing-xi-de-tui-dao-si-lu-by-mrsate/
这个题解写的比较好理解
*/
func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n2 := dp[a] * 2
		n3 := dp[b] * 3
		n5 := dp[c] * 5
		if n2 <= n3 && n2 <= n5 {
			dp[i] = n2
		} else if n3 <= n2 && n3 <= n5 {
			dp[i] = n3
		} else {
			dp[i] = n5
		}
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
