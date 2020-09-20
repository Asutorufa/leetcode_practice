package main

import (
	"container/heap"
	"fmt"
)

/**
如何得到一个数据流中的中位数？
如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，
[2,3,4] 的中位数是 3
[2,3] 的中位数是 (2 + 3) / 2 = 2.5
设计一个支持以下两种操作的数据结构：
    void addNum(int num) - 从数据流中添加一个整数到数据结构中。
    double findMedian() - 返回目前所有元素的中位数。

示例 1：
输入：
["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]
输出：[null,null,null,1.50000,null,2.00000]

示例 2：
输入：
["MedianFinder","addNum","findMedian","addNum","findMedian"]
[[],[2],[],[3],[]]
输出：[null,null,2.00000,null,2.50000]

限制：
    最多会对 addNum、findMedia进行 50000 次调用。
*/

type MedianFinder struct {
	nums []int
	even bool

	lowerHeap  *heapI
	higherHeap *heapH
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	x := &heapI{}
	y := &heapH{}
	heap.Init(x)
	heap.Init(y)
	return MedianFinder{
		lowerHeap:  x,
		higherHeap: y,
	}
}

// 大小端堆实现
/**
A -> 小端堆 B -> 大端堆
算法流程：

    设元素总数为 N=m+n ，其中 m 和 n 分别为 A 和 B 中的元素个数。

addNum(num) 函数：

    当 m=n（即 N 为 偶数）：需向 A 添加一个元素。实现方法：将新元素 num 插入至 B ，再将 B 堆顶元素插入至 A ；
    当 m!=n（即 N 为 奇数）：需向 B 添加一个元素。实现方法：将新元素 num 插入至 A ，再将 A 堆顶元素插入至 B ；

	假设插入数字 num 遇到情况 1. 。由于 num 可能属于 “较小的一半” （即属于 B ），
	因此不能将 nums 直接插入至 A 。而应先将 num 插入至 B ，
	再将 B堆顶元素插入至 A 。这样就可以始终保持 A 保存较大一半、 B 保存较小一半。

findMedian() 函数：
    当 m=n（ N 为 偶数）：则中位数为 ( A 的堆顶元素 + B 的堆顶元素 )/2。
    当 m!=n（ N 为 奇数）：则中位数为 A 的堆顶元素。
*/
func (this *MedianFinder) AddNum(num int) {
	if this.lowerHeap.Len() == this.higherHeap.Len() {
		heap.Push(this.higherHeap, num)
		x := heap.Pop(this.higherHeap).(int)
		// fmt.Println(this.higherHeap)
		// fmt.Println("pop", x)
		heap.Push(this.lowerHeap, x)
		return
	}

	heap.Push(this.lowerHeap, num)
	x := heap.Pop(this.lowerHeap).(int)
	// fmt.Println("pop", x)
	heap.Push(this.higherHeap, x)
}

func (this *MedianFinder) FindMedian() float64 {
	if this.higherHeap.Len() != this.lowerHeap.Len() {
		x := heap.Pop(this.lowerHeap).(int)
		heap.Push(this.lowerHeap, x)
		return float64(x)
	}

	x := heap.Pop(this.higherHeap).(int)
	heap.Push(this.higherHeap, x)
	y := heap.Pop(this.lowerHeap).(int)
	heap.Push(this.lowerHeap, y)
	return (float64(x) + float64(y)) / 2
}

// 超出时间限制
func (this *MedianFinder) AddNum2(num int) {
	l := len(this.nums)
	if l == 0 {
		this.nums = append(this.nums, num)
		// fmt.Println(this.nums)
		return
	}
	if l == 1 {
		if num > this.nums[0] {
			this.nums = append(this.nums, num)
			// fmt.Println(this.nums)
			return
		}
		this.nums = append([]int{num}, this.nums[0])
		// fmt.Println(this.nums)
		return
	}
	for i := 0; i < l; i++ {
		if i == 0 && num < this.nums[i] {
			this.nums = append([]int{num}, this.nums...)
			break
		}
		if i == l-1 && num >= this.nums[i] {
			this.nums = append(this.nums, num)
			break
		}

		if num >= this.nums[i] && num < this.nums[i+1] {
			i++
			var pre = make([]int, i)
			var last = make([]int, l-i)
			copy(pre, this.nums[:i])
			copy(last, this.nums[i:])
			// fmt.Println(pre, last)
			this.nums = append(pre, num)
			this.nums = append(this.nums, last...)
			break
		}
	}
	// fmt.Println(this.nums)
}

func (this *MedianFinder) FindMedian2() float64 {
	l := len(this.nums)
	if l == 0 {
		return 0
	}
	if l%2 != 0 {
		return float64(this.nums[l/2])
	}
	x := l/2 - 1
	return (float64(this.nums[x]) + float64(this.nums[x+1])) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

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

type heapH []int

func (h heapH) Len() int {
	return len(h)
}

func (h heapH) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h heapH) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapH) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heapH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	obj := Constructor()
	obj.AddNum(-1)
	fmt.Println(obj.FindMedian())
	obj.AddNum(-2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(-3)
	fmt.Println(obj.FindMedian())
	obj.AddNum(-4)
	fmt.Println(obj.FindMedian())
	obj.AddNum(-5)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())

	fmt.Println("\n[null,null,null,1.50000,null,2.00000]")
	obj = Constructor()
	obj.AddNum(1)
	fmt.Println(obj.FindMedian())
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())

	fmt.Println("\n[[],[6],[],[10],[],[2],[],[6],[],[5],[],[0],[],[6],[],[3],[],[1],[],[0],[],[0],[]]")
	obj = Constructor()
	obj.AddNum(6)
	fmt.Println(obj.FindMedian())
	obj.AddNum(10)
	fmt.Println(obj.FindMedian())
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(6)
	fmt.Println(obj.FindMedian())
	obj.AddNum(5)
	fmt.Println(obj.FindMedian())
	obj.AddNum(0)
	fmt.Println(obj.FindMedian())
	obj.AddNum(6)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
	obj.AddNum(1)
	fmt.Println(obj.FindMedian())
	obj.AddNum(0)
	fmt.Println(obj.FindMedian())
	obj.AddNum(0)
	fmt.Println(obj.FindMedian())

	fmt.Println("\nheap higher test")
	higherTest := &heapH{}
	heap.Init(higherTest)
	heap.Push(higherTest, 1)
	heap.Push(higherTest, 2)
	heap.Push(higherTest, 3)
	heap.Push(higherTest, 4)
	fmt.Println(heap.Pop(higherTest))
	fmt.Println(heap.Pop(higherTest))
	fmt.Println(heap.Pop(higherTest))

	fmt.Println("\nheap lower test")
	lowerTest := &heapI{}
	heap.Init(lowerTest)
	heap.Push(lowerTest, 1)
	heap.Push(lowerTest, 2)
	heap.Push(lowerTest, 3)
	fmt.Println(heap.Pop(lowerTest))
	fmt.Println(heap.Pop(lowerTest))
}
