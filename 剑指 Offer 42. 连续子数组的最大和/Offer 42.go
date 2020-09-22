package main

import "fmt"

/**
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
要求时间复杂度为O(n)。

示例1:
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

提示：
    1 <= arr.length <= 10^5
    -100 <= arr[i] <= 100
*/

// 动态规划
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	index := 0
	i := make([]int, l)
	i[index] = nums[index]
	max := i[0]
	index++

	for ; index < l; index++ {
		if i[index-1] > 0 {
			i[index] = nums[index] + i[index-1]
		} else {
			i[index] = nums[index]
		}
		if i[index] > max {
			max = i[index]
		}
	}
	// fmt.Println(i)
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
