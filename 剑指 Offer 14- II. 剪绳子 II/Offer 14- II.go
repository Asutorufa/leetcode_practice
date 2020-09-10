package main

import (
	"fmt"
	"math/big"
)

/**
给你一根长度为 n 的绳子，
请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
每段绳子的长度记为 k[0],k[1]...k[m - 1] 。
请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1

示例 2:
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36

提示：

    2 <= n <= 1000
*/

// 与14-I相同, 但是需要注意越界问题
func cuttingRope(n int) int {
	if n <= 3 { // 稍加推导即可得出
		return n - 1
	}
	dp := make([]*big.Int, n+1)
	dp[0] = big.NewInt(0)
	dp[1] = big.NewInt(1)
	dp[2] = big.NewInt(2)
	dp[3] = big.NewInt(3) // 此处与上面不同是因为这里是作为乘数, 且自己比分段大, 所以直接使用自己的值就行了

	for i := 4; i <= n; i++ { // 转化为每个小段长度乘积的问题
		// fmt.Println("当前长度:", i)
		max := big.NewInt(0)
		for j := 1; j <= i/2; j++ { // 除以2是因为, 类似(2,3)和(3,2)这种相同, 所以应当除以2
			// fmt.Println("长度为", j, "的最大分段乘积(", dp[j], ")乘以长度为", i-j, "的最大分段乘积(", dp[i-j], ")")
			tmp := big.NewInt(0).Mul(dp[j], dp[i-j])
			if tmp.Cmp(max) == 1 {
				max = tmp
			}
		}
		// fmt.Println("长度为", i, "的最大分段乘积为:", max)
		dp[i] = max
	}
	return int(big.NewInt(0).Mod(dp[n], big.NewInt(1000000007)).Int64())
}

func main() {
	fmt.Println(cuttingRope(2))    // 1
	fmt.Println(cuttingRope(10))   // 36
	fmt.Println(cuttingRope(20))   // 1458
	fmt.Println(cuttingRope(120))  // 953271190
	fmt.Println(cuttingRope(127))  // 439254203
	fmt.Println(cuttingRope(999))  // 965709895
	fmt.Println(cuttingRope(2000)) // 545483727
}
