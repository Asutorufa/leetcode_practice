package main

import "fmt"

/**
在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

示例 1：
输入：nums = [3,4,3,3]
输出：4

示例 2：
输入：nums = [9,1,7,9,7,9,7]
输出：1

限制：
    1 <= nums.length <= 10000
    1 <= nums[i] < 2^31
*/

// 通过统计数二进制的1的个数,然后用3求余,剩下的就是那一个数的二进制
func singleNumber(nums []int) int {
	count := make([]int, 32)
	for index := range nums {
		for i := 0; i < 32; i++ {
			count[i] += nums[index] & 1
			nums[index] >>= 1
		}
	}
	// fmt.Println(count)
	res := 0
	for i := 0; i < 32; i++ {
		res <<= 1
		res |= count[31-i] % 3
	}
	return res
}

/**
有限状态自动机 强(数字逻辑,计算机组成原理,编译原理相关)
 状态转移表
  0 ---> 1 ---> 2 -+
  ^			       |
  |			       |
  +----------------+
  输入为 1 的话就转移到下一个状态, 输入为 0 的话就保持当前状态
  三个状态用二进制表示需要两位数
  00 ---> 01 ---> 10 -+
  ^                   |
  |                   |
  +-------------------+


初始：0 0
1个1：1 0
2个1：0 1
3个1：0 0

一下公式由来题解:
	[](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/solution/zhuang-tai-ji-jie-jue-ci-wen-ti-xiang-jie-shu-zi-d/)
*/
func singleNumber2(nums []int) int {
	one, two := 0, 0
	for index := range nums {
		one = one ^ nums[index] & ^two // ^two 按位取反
		two = two ^ nums[index] & ^one // ^one 按位取反
	}
	return one
}

func main() {
	fmt.Println(singleNumber([]int{9, 1, 7, 9, 7, 9, 7}))
	fmt.Println(singleNumber([]int{3, 4, 3, 3}))
	fmt.Println(singleNumber2([]int{9, 1, 7, 9, 7, 9, 7}))
	fmt.Println(singleNumber2([]int{3, 4, 3, 3}))
}
