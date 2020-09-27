package main

import "fmt"

/**
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

示例 1：
输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]

示例 2：
输入：nums = [1,2,10,4,1,4,3,3]
输出：[2,10] 或 [10,2]

限制：
    2 <= nums.length <= 10000
*/

/**
异或的性质 一开始还没想清楚,我太笨了....
>>> bin(0b010 ^ 0b101)
'0b111'
>>> bin(0b011 ^ 0b101)
'0b110'
>>> bin(0b111 ^ 0b101)
'0b10'
*/

func singleNumbers(nums []int) []int {
	//k用来计算nums的异或和
	k := 0
	// 计算异或和 并存到k
	// e.g. [2,4,2,3,3,6] 异或和：(2^2)^(3^3)^(4^6)=2=010
	for index := range nums {
		k ^= nums[index]
	}
	// fmt.Println(k)

	//设置mask为1，则二进制为0001
	// mask是一个二进制数，且其中只有一位是1，其他位全是0，比如000010，
	// 表示我们用倒数第二位作为分组标准，倒数第二位是0的数字分到一组，倒数第二位是1的分到另一组
	mask := 1
	// & operator只有1&1时等于1 其余等于0
	// 用上面的e.g. 4和6的二进制是不同的 我们从右到左找到第一个不同的位就可以分组 4=0100 6=0110
	// 根据e.g. 010 & 001 = 000 = 0则 mask=010
	// 010 & 010 != 0 所以mask=010
	// 之后就可以用mask来将数组里的两个数分区分开
	for mask&k == 0 {
		mask <<= 1
		// fmt.Printf("%b %b\n", mask, k)
	}
	// fmt.Println(mask)

	a, b := 0, 0
	for index := range nums {
		//根据&是否为0区分将两个数字分区，并分别求异或和
		if nums[index]&mask == 0 {
			a ^= nums[index]
		} else {
			b ^= nums[index]
		}
	}
	// fmt.Println(a, b)
	return []int{a, b}
}

func main() {
	fmt.Println(singleNumbers([]int{1, 2, 10, 4, 1, 4, 3, 3}))
	fmt.Println(singleNumbers([]int{1, 55, 249, 4, 1, 4, 3, 3}))
}
