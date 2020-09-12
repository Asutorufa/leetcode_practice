package main

import "fmt"

/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

示例：
输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。

提示：
    1 <= nums.length <= 50000
    1 <= nums[i] <= 10000
*/

// 前后双指针
func exchange(nums []int) []int {
	start := 0
	end := len(nums) - 1
	for start < end {
		for start < end && nums[start]%2 != 0 {
			start++
		}
		for start < end && nums[end]%2 == 0 {
			end--
		}
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
	}
	return nums
}

// 全部循环 用两个数组存 最后合并
func exchange2(nums []int) []int {
	var odd []int
	var even []int
	for index := range nums {
		if nums[index]%2 == 0 {
			even = append(even, nums[index])
			continue
		}
		odd = append(odd, nums[index])
	}
	return append(odd, even...)
}

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
	fmt.Println(exchange2([]int{1, 2, 3, 4}))
}
