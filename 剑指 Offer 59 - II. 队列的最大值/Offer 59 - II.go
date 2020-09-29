package main

import (
	"fmt"
)

/**
请定义一个队列并实现函数 max_value 得到队列里的最大值，
要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：
输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]

示例 2：
输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]

限制：
    1 <= push_back,pop_front,max_value的总操作数 <= 10000
    1 <= value <= 10^5
*/

// 单调队列 + 正常队列实现 ... 比较复杂
// 单调队列与 59 - I 中的相同, 其实想想就能很好理解
type MaxQueue struct {
	queue     *queue     // 正常队列
	monoQueue *monoQueue // 单调队列
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue:     &queue{},
		monoQueue: &monoQueue{},
	}
}

func (this *MaxQueue) Max_value() int {
	return this.monoQueue.front()
}

func (this *MaxQueue) Push_back(value int) {
	this.queue.push(value)
	this.monoQueue.push(value)
}

func (this *MaxQueue) Pop_front() int {
	x := this.queue.pop()
	if this.monoQueue.front() == x { // 如果当前队列出的值等于最大值 就从单调队列中弹出
		this.monoQueue.pop(x)
	}
	return x
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

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
	return -1
}

type queue struct {
	que []int
}

func (q *queue) push(a int) {
	q.que = append(q.que, a)
}

func (q *queue) pop() int {
	if q.size() <= 0 {
		return -1
	}
	x := q.que[0]
	q.que = q.que[1:]
	return x
}

func (q *queue) peek() int {
	if q.size() <= 0 {
		return -1
	}
	return q.que[0]
}

func (q *queue) size() int {
	return len(q.que)
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Pop_front())
	fmt.Println(obj.Max_value())
	obj.Push_back(1)
	obj.Push_back(2)
	fmt.Println(obj.Max_value())
	fmt.Println(obj.Pop_front())
	fmt.Println(obj.Max_value())
	fmt.Println(obj.Pop_front())
	fmt.Println()
	obj = Constructor()
	obj.Push_back(94)
	obj.Push_back(16)
	obj.Push_back(89)
	fmt.Println(obj.Pop_front())
	obj.Push_back(22)
	fmt.Println(obj.Pop_front())
}
