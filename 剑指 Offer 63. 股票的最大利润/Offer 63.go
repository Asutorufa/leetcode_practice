package main

import "fmt"

/**
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

示例 1:
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

示例 2:
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

限制：
0 <= 数组长度 <= 10^5
*/

// 自己写的 还算好理解... 跟最佳题解差不多就是了...
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	sum, max, monoQueue := 0, prices[0], []int{prices[0]}
	for i := 1; i < l; i++ {
		if prices[i] < monoQueue[0] {
			monoQueue = []int{}
			max = prices[i]
		}
		monoQueue = append(monoQueue, prices[i])
		// fmt.Println(monoQueue, max, sum)
		lt := len(monoQueue)
		if lt > 1 && monoQueue[lt-1] > max {
			x := monoQueue[lt-1] - monoQueue[0]
			if x > sum {
				sum = x
			}
			max = monoQueue[lt-1]
		}
	}
	return sum
}

// 别人的题解  确实写的比较好
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, profit := prices[0], 0

	for index := range prices {
		if prices[index] < min {
			min = prices[index]
		}
		tmp := prices[index] - min
		if profit < tmp {
			profit = tmp
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 6, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{3, 2, 6, 5, 0, 3}))
	fmt.Println()
	fmt.Println(maxProfit2([]int{7, 1, 5, 6, 3, 6, 4}))
	fmt.Println(maxProfit2([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit2([]int{3, 2, 6, 5, 0, 3}))
}
