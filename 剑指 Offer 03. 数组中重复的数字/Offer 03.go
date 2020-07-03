package main

import "log"

/*
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3



限制：

2 <= n <= 100000
*/

// 利用map存储 占用内存大
func findRepeatNumber(nums []int) int {
	tmp := make(map[int]bool)
	for index := range nums {
		if tmp[nums[index]] == true {
			return nums[index]
		}
		tmp[nums[index]] = true
	}
	return 0
}

// 利用int数组存储 占用内存小
func findRepeatNumber2(nums []int) int {
	tmp := make([]int, len(nums))
	for index := range nums {
		if tmp[nums[index]] < 0 {
			return nums[index]
		}
		tmp[nums[index]] = -1
	}
	return 0
}

func main() {
	log.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
	log.Println(findRepeatNumber2([]int{2, 3, 1, 0, 2, 5, 3}))
}
