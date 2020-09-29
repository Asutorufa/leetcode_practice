package main

import (
	"fmt"
	"sort"
)

/**
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。
2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。
A 不能视为 14。

示例 1:
输入: [1,2,3,4,5]
输出: True

示例 2:
输入: [0,0,1,2,5]
输出: True

限制：
数组长度为 5
数组的数取值为 [0, 13] .
*/

func isStraight(nums []int) bool {
	sort.Ints(nums)
	// fmt.Println(nums)
	zero := 0
	for index := range nums {
		// fmt.Println(nums[index])
		if nums[index] == 0 {
			zero++
			continue
		}
		if index == 0 || nums[index-1] == 0 {
			continue
		}
		x := nums[index] - nums[index-1]
		// fmt.Println("x", x, nums[index], nums[index-1])
		if x <= 0 || zero < x-1 {
			return false
		}
		zero -= x - 1
	}
	return true
}

// 注意题目中只抽5张牌 因此只要最大值 - 最小值 > 5 即可构成
func isStraight2(nums []int) bool {
	min, max := 14, 0
	rep := map[int]int{}
	for index := range nums {
		if nums[index] == 0 {
			continue
		}
		if nums[index] > max {
			max = nums[index]
		}
		if nums[index] < min {
			min = nums[index]
		}
		if _, ok := rep[nums[index]]; ok {
			return false
		}
		rep[nums[index]] = nums[index]
	}
	return max-min < 5
}

func main() {
	fmt.Println(isStraight([]int{0, 0, 1, 2, 5}))
	fmt.Println(isStraight([]int{0, 1, 2, 5, 6}))
	fmt.Println(isStraight([]int{0, 0, 1, 2, 6}))
	fmt.Println(isStraight([]int{0, 0, 8, 5, 4}))
	fmt.Println(isStraight([]int{0, 0, 2, 2, 5}))
	fmt.Println()
	fmt.Println(isStraight2([]int{0, 0, 1, 2, 5}))
	fmt.Println(isStraight2([]int{0, 1, 2, 5, 6}))
	fmt.Println(isStraight2([]int{0, 0, 1, 2, 6}))
	fmt.Println(isStraight2([]int{0, 0, 8, 5, 4}))
	fmt.Println(isStraight2([]int{0, 0, 2, 2, 5}))
}
