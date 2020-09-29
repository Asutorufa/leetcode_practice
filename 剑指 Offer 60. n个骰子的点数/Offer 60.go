package main

import (
	"fmt"
	"math"
)

/**
把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

示例 1:
输入: 1
输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]

示例 2:
输入: 2
输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]

限制：
1 <= n <= 11
*/

// 动态优化
// 相对比较好理解 用上一个状态(就是比当前少一个骰子的状态)的每个值加一遍1-6每个值即可得到当前状态
// 			   或者也可说是当前要求的骰子的和减去1-6,每个数减一遍就是上个状态的次数值
func twoSum(n int) []float64 {
	dp := make([][]int, n+1) // 第一层 -> 表示几个骰子的值
	for index := range dp {
		dp[index] = make([]int, 6*(index+1)+1) // 第二层, 每个和的出现的次数次数
	}

	for i := 1; i <= 6; i++ {
		dp[1][i] = 1 // 第一层, 也就是当为一个骰子时, 可知 1-6 出现的次数都是1
	}

	for i := 2; i <= n; i++ { // 骰子的个数
		for j := i; j <= 6*i; j++ { // 总和的个数
			for cur := 1; cur <= 6; cur++ { // 遍历减去 1-6, 得到上个状态的次数和并相加
				if j-cur <= 0 { // 不存在骰子数小于等于0
					break
				}
				dp[i][j] += dp[i-1][j-cur]
			}
		}
	}
	all := math.Pow(6, float64(n)) // 总共会出现的情况
	res := []float64{}

	for i := n; i <= 6*n; i++ {
		res = append(res, float64(dp[n][i])/all) // 次数除以总共情况, 可得概率
	}
	return res
}

// 存储优化版
func twoSum2(n int) []float64 {
	dp := make([]int, 6*n+1)
	for i := 1; i <= 6; i++ {
		dp[i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 6 * i; j >= i; j-- { // 倒着遍历 这样上一个状态的数据就不会被覆盖了
			dp[j] = 0 // 把之前的数据置零
			for cur := 1; cur <= 6; cur++ {
				if j-cur < i-1 { // 上一个状态的最小值 比如: 6个骰子的最小值就是6
					break
				}
				dp[j] += dp[j-cur]
			}
		}
	}
	all := math.Pow(6, float64(n))
	res := []float64{}

	for i := n; i <= 6*n; i++ {
		res = append(res, float64(dp[i])/all)
	}
	return res
}

func main() {
	fmt.Println(twoSum(1))
	fmt.Println(twoSum(2))
	fmt.Println(twoSum(3))
	fmt.Println(twoSum2(1))
	fmt.Println(twoSum2(2))
	fmt.Println(twoSum2(3))
}
