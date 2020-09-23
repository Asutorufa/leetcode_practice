package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

示例 1:
输入: [10,2]
输出: "102"

示例 2:
输入: [3,30,34,5,9]
输出: "3033459"

提示:
    0 < nums.length <= 100

说明:
    输出结果可能非常大，所以你需要返回一个字符串而不是整数
    拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
*/

/**
此题求拼接起来的 “最小数字” ，本质上是一个排序问题。
排序判断规则： 设 nums 任意两数字的字符串格式 x 和 y ，则
    若拼接字符串 x+y>y+x ，则 m>n ；
    反之，若 x+y<y+x ，则 n<m ；

证明:
字符串 xy < yx , yz < zy ，需证明 xz < zx 一定成立。

设十进制数 x, y, z 分别有 a, b, c 位，则有：
（左边是字符串拼接，右边是十进制数计算，两者等价）
xy = x * 10^b + y
yx = y * 10^a + x

则 xy < yx 可转化为：
x * 10^b + y < y * 10^a + x
x (10^b - 1) < y (10^a - 1)
x / (10^a - 1) < y / (10^b - 1)     ①

同理， 可将 yz < zy 转化为：
y / (10^b - 1) < z / (10^c - 1)     ②

将 ① ② 合并，整理得：
x / (10^a - 1) < y / (10^b - 1) < z / (10^c - 1)
x / (10^a - 1) < z / (10^c - 1)
x (10^c - 1) < z (10^a - 1)
x * 10^c + z < z * 10^a + x
∴  可推出 xz < zx ，传递性证毕
*/

func minNumber(nums []int) string {
	numStr := []string{}
	for index := range nums {
		numStr = append(numStr, strconv.Itoa(nums[index]))
	}
	fastSort(0, len(numStr)-1, numStr)
	// fmt.Println(numStr)
	return strings.Join(numStr, "")
}

// 稍微改一下 快速排序 原版在剑指 Offer 40
func fastSort(left, right int, numStr []string) {
	if left > right {
		return
	}
	low, high := left, right
	target := numStr[left]
	for left < right {
		for left < right && numStr[right]+target >= target+numStr[right] { // == numStr[right] >= target
			right--
		}
		for left < right && numStr[left]+target <= target+numStr[left] { // == numStr[left] <= target
			left++
		}
		numStr[left], numStr[right] = numStr[right], numStr[left]
	}
	numStr[left], numStr[low] = numStr[low], numStr[left]
	fastSort(low, left-1, numStr)
	fastSort(right+1, high, numStr)
}

func main() {
	fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
}
