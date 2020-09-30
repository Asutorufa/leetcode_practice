package main

import "fmt"

/**
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，
其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

示例:
输入: [1,2,3,4,5]
输出: [120,60,40,30,24]

提示：
    所有元素乘积之和不会溢出 32 位整数
    a.length <= 100000
*/

// 上下三角计算方法
/**
如 1 2 3 4 5

* 2 3 4 5
1 * 3 4 5
1 2 * 4 5
1 2 3 * 5
1 2 3 4 *

先计算两个三角
	左下角三角过程:
		*			a[0] = 1
		1 * 1 = 1	a[1] = a[1-1] * nums[1-1] = 1 * 1
		1 * 2 = 2   a[2] = a[2-1] * nums[2-1] = 1 * 2
		2 * 3 = 6   a[3] = a[3-1] * nums[3-1] = 2 * 3
		6 * 4 = 24  a[4] = a[4-1] * nums[4-1] = 6 * 4
	右下角计算过程:
		*			  b[len-1] = 1								 = b[4]
		5 * 1  = 5    b[len-2] = b[len-1] * nums[len-1] = 1 * 5  = b[3]
		5 * 4  = 20   b[len-3] = b[len-2] * nums[len-2] = 5 * 4  = b[2]
		20 * 3 = 60   b[len-4] = b[len-3] * nums[len-3] = 20 * 3 = b[1]
		60 * 2 = 120  b[len-5] = b[len-4] * nums[len-5] = 60 * 2 = b[0]
	合并:
		sum[0] = a[0] * b[0] = 1  * 120 = 120
		sum[1] = a[1] * b[1] = 1  * 60  = 60
		sum[2] = a[2] * b[2] = 2  * 20  = 40
		sum[3] = a[3] * b[3] = 6  * 5   = 30
		sum[4] = a[4] * b[4] = 24 * 1   = 24
*/
func constructArr(a []int) []int {
	l := len(a)
	if l == 0 {
		return make([]int, 0)
	}
	b := make([]int, l)
	b[0] = 1
	c := make([]int, l)
	c[l-1] = 1
	for index := 1; index < l; index++ {
		b[index] = b[index-1] * a[index-1]
		c[l-index-1] = c[l-index] * a[l-index]
	}
	for index := l - 2; index >= 0; index-- {
		b[index] *= c[index]
	}
	// fmt.Println(b, c)
	return b
}

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4}))
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}
