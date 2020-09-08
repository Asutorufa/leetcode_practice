package main

import (
	"fmt"
)

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：

输入：[3,4,5,1,2]
输出：1

示例 2：

输入：[2,2,2,0,1]
输出：0
*/

/*
	Brute Force
*/
func minArray(numbers []int) int {
	for i := len(numbers) - 1; i >= 0; i-- {
		if i < len(numbers)-1 && numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
		if i == 0 {
			return numbers[0]
		}
	}
	return 0
}

/*
二分查找
注: m=mid i=low j=high
算法流程：

    初始化： 声明 i, j 双指针分别指向 nums 数组左右两端；
    循环二分： 设 m=(i+j)/2为每次二分的中点（ "/" 代表向下取整除法，因此恒有 i≤m<j ），可分为以下三种情况：
        当 nums[m]>nums[j] 时： m 一定在 左排序数组 中，即旋转点 x 一定在 [m+1,j] 闭区间内，因此执行 i=m+1；
        当 nums[m]<nums[j] 时： m 一定在 右排序数组 中，即旋转点 x 一定在[i,m] 闭区间内，因此执行 j=m；
        当 nums[m]=nums[j] 时： 无法判断 m 在哪个排序数组中，即无法判断旋转点 x 在 [i,m] 还是 [m+1,j] 区间中。解决方案： 执行 j=j−1 缩小判断范围，分析见下文。
    返回值： 当 i=j 时跳出二分循环，并返回 旋转点的值 nums[i] 即可。
*/
func minArray2(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		mid := low + (high-low)/2
		if numbers[mid] < numbers[high] {
			high = mid
		} else if numbers[mid] > numbers[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return numbers[low]
}

func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
	fmt.Println(minArray([]int{1, 3, 5}))
	fmt.Println(minArray2([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray2([]int{2, 2, 2, 0, 1}))
	fmt.Println(minArray2([]int{1, 3, 5}))
}
