package main

import "fmt"

/**
统计一个数字在排序数组中出现的次数。

示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: 2

示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: 0

限制：
0 <= 数组长度 <= 50000
*/

// 两种差不多 甚至第二种测试结果还没我第一个暴力写的好看...
func search(nums []int, target int) int {
	res := 0

	for index := range nums{
		if res != 0 && nums[index] != target{
			break
		}
		if nums[index] == target{
			res++
		}
	}
	return res
}

func search2(nums []int,target int) int{
	l := len(nums)
	i,j := 0,l-1
	for ;i<l;i++{
		if nums[i] == target{
			i--
			break
		}
	}

	if i >= l{
		return 0
	}

	for ;j>=0;j--{
		if nums[j] == target{
			j++
			break
		}
	}

	return j-(i+1)
}

func main(){
	fmt.Println(search([]int{5,7,7,8,8,10},8))
	fmt.Println(search([]int{5,7,7,8,8,10},6))
	fmt.Println(search2([]int{5,7,7,8,8,10},8))
	fmt.Println(search2([]int{5,7,7,8,8,10},6))
}