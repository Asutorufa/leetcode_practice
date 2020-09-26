package main

import "fmt"

/**
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

示例 1:
输入: [0,1,3]
输出: 2

示例 2:
输入: [0,1,2,3,4,5,6,7,9]
输出: 8

限制：
1 <= 数组长度 <= 10000
*/

// 二分法 更快
func missingNumber(nums []int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		m := (i + j) / 2
		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}

// 自己写的 双指针... 不行.. 虽然测试出来结果差不多
func missingNumber2(nums []int) int {
	l := len(nums)
	if nums[0] != 0 {
		return 0
	}
	i, j := 0, l-1
	for i < l && j >= 0 {
		if i <= 0 && j >= l-1 {
			i++
			j--
			continue
		}
		if nums[i]-nums[i-1] > 1 {
			return nums[i] - 1
		}
		if nums[j+1]-nums[j] > 1 {
			return nums[j+1] - 1
		}
		j--
		i++
	}
	return nums[l-1] + 1
}

func main() {
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
	fmt.Println(missingNumber([]int{0, 1, 3}))
	fmt.Println(missingNumber([]int{0}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{0, 2}))
	fmt.Println(missingNumber([]int{1}))
	fmt.Println()
	fmt.Println(missingNumber2([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
	fmt.Println(missingNumber2([]int{0, 1, 3}))
	fmt.Println(missingNumber2([]int{0}))
	fmt.Println(missingNumber2([]int{0, 1}))
	fmt.Println(missingNumber2([]int{0, 2}))
	fmt.Println(missingNumber2([]int{1}))

}
