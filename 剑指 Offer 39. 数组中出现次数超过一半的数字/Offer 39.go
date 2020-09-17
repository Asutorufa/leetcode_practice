package main

import "fmt"

/**
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:
输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2

限制：
1 <= 数组长度 <= 50000
*/

// my solution: use map
func majorityElement(nums []int) int {
	s := map[int]int{}
	for _, x := range nums {
		s[x]++
		if s[x] > len(nums)/2 {
			return x
		}
	}
	return 0
}

// 投票法找众数, 很神奇
func majorityElement2(nums []int) int {
	x, vote := 0, 0 // x: 被投数字, vote: 票数

	for _, y := range nums {
		// fmt.Println("num", x, "vote", vote)
		if vote == 0 { // 若投票数為0 就将当前的数字设置为被投的数字
			x = y
			vote++
			continue
		}

		if x != y { // 如果当前数字与被投数字不同 票数-1
			vote--
			continue
		}
		if x == y { // 若当前数字与被投数字相同 票数+1
			vote++
		}
	}
	// fmt.Println("num", x, "vote", vote)

	if vote > 0 { // 判断票数是否大于0
		return x
	}
	return 0
}

func main() {
	fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
	fmt.Println(majorityElement2([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
