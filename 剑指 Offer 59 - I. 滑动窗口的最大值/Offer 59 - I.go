package main

import "fmt"

/**
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:
  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

提示：
你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
*/

// 暴力解法
func maxSlidingWindow(nums []int, k int) []int {
	i, j, l := 0, k, len(nums)
	if l == 0 {
		return []int{}
	}
	res := []int{}
	for j <= l {
		res = append(res, getMax(nums[i:j]))
		i++
		j++
	}
	return res
}

func getMax(a []int) (res int) {
	res = a[0]
	for index := range a {
		if a[index] > res {
			res = a[index]
		}
	}
	return
}

// 单调队列解法 属实强
// 题解: https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/solution/dan-diao-dui-lie-xiang-jie-by-carlsun-2/
func maxSlidingWindow2(nums []int, k int) []int {
	i, j := 1, k
	l := len(nums)
	if l == 0 {
		return []int{}
	}
	res := []int{}
	x := &monoQueue{}

	for z := 0; z < k; z++ {
		x.push(nums[z])
	}
	res = append(res, x.front())

	for j < l {
		// fmt.Println(x.queue)
		x.pop(nums[i-1])
		x.push(nums[j])
		res = append(res, x.front())
		i++
		j++
	}
	return res
}

type monoQueue struct {
	queue []int
}

func (m *monoQueue) push(a int) {
	if len(m.queue) == 0 {
		m.queue = append(m.queue, a)
		return
	}
	// if len(m.queue) != 0 {
	// fmt.Println(m.queue, m.queue[0], a, m.queue[0] < a)
	// }
	for len(m.queue) != 0 && m.queue[len(m.queue)-1] < a {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, a)
}

func (m *monoQueue) pop(a int) {
	if len(m.queue) != 0 && m.queue[0] == a {
		m.queue = m.queue[1:]
	}
}

func (m *monoQueue) front() int {
	if len(m.queue) != 0 {
		return m.queue[0]
	}
	return 0
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow2([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow2([]int{1}, 1))
	fmt.Println(maxSlidingWindow2([]int{}, 0))
	fmt.Println(maxSlidingWindow2([]int{1, 3, 1, 2, 0, 5}, 3))

	fmt.Println("\nQueue Test")
	x := &monoQueue{}
	x.push(1)
	fmt.Println(x.queue, x.front())
	x.push(2)
	fmt.Println(x.queue, x.front())
	x.push(3)
	fmt.Println(x.queue, x.front())
	x.push(-1)
	fmt.Println(x.queue, x.front())
	x.pop(3)
	fmt.Println(x.queue, x.front())
}
