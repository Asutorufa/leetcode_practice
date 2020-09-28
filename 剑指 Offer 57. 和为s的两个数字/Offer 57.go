package main

import "fmt"

// 与 leetcode 第一题, 稍微不同, leetcode 第一题不是排好序的数组

/**
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
如果有多对数字的和等于s，则输出任意一对即可。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[2,7] 或者 [7,2]

示例 2：
输入：nums = [10,26,30,31,47,60], target = 40
输出：[10,30] 或者 [30,10]

限制：
    1 <= nums.length <= 10^5
    1 <= nums[i] <= 10^6
*/

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		s := nums[i] + nums[j]
		if s > target {
			j--
		} else if s < target {
			i++
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
