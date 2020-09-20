package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/**
输入整数数组 arr ，找出其中最小的 k 个数。
例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例 1：
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]

示例 2：
输入：arr = [0,1,2,1], k = 1
输出：[0]

限制：
    0 <= k <= arr.length <= 10000
    0 <= arr[i] <= 10000
*/

func getLeastNumbers2(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

/**
使用快速排序
*/
func getLeastNumbers3(arr []int, k int) (res []int) {
	return fastSort(arr)[:k]
}

func fastSort(arr []int) []int {
	// fmt.Println(arr)
	l := len(arr)
	if l == 0 || l == 1 {
		return arr
	}
	base := arr[0]
	var i int = 0
	var j int = l - 1

	for {
		for {
			if i <= l-1 && arr[i] <= base {
				i++
				continue
			}
			break
		}

		for {
			if j >= 0 && arr[j] > base {
				j--
				continue
			}
			break
		}

		if i >= j {
			break
		}
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	}

	var index int
	if i > j { // 选择交换 和 分割位置
		index = j
	} else {
		index = i
	}
	arr[0] = arr[index]
	arr[index] = base

	var pre []int
	var last []int
	if index+1 < l {
		last = arr[index+1:]
	}
	pre = arr[:index]

	arr = append(fastSort(pre), base)
	arr = append(arr, fastSort(last)...)
	return arr
}

// 利用堆实现 小端堆
func getLeastNumbers(arr []int, k int) (res []int) {
	x := &heapI{}
	*x = arr
	heap.Init(x)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(x).(int))
	}

	return res
}

type heapI []int

func (h heapI) Len() int {
	return len(h)
}

func (h heapI) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h heapI) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapI) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heapI) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println(getLeastNumbers([]int{56, 7, 7, 4, 54, 4}, 3))
	fmt.Println(getLeastNumbers2([]int{56, 7, 7, 4, 54, 4}, 3))
	fmt.Println(getLeastNumbers3([]int{56, 7, 7, 4, 54, 4}, 3))

	fmt.Println("\nfastOpen Test ")
	fmt.Println(fastSort([]int{2, 1, 1, 2}))
	fmt.Println(fastSort([]int{3, 1, 1, 2, 5}))
	fmt.Println(fastSort([]int{3, 1, 1, 2, 2, 5}))
	fmt.Println(fastSort([]int{3, 1, 1, 2, 98, 97, 2, 5, 55, 63}))
	fmt.Println(fastSort([]int{3, 1, 1, 2, 96, 98, 97, 2, 5, 55, 63}))
}
