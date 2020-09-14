package main

import (
	"fmt"
)

/**
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
调用 min、push 及 pop 的时间复杂度都是 O(1)。

示例:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

提示：
    各函数的调用总次数不超过 20000 次
*/

/**
双栈实现, 想不到...
*/

// MinStack min stack
type MinStack struct {
	stack []int
	order []int
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{}
}

// Push push
func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.order) == 0 {
		this.order = append(this.order, x)
	} else {
		if x <= this.order[len(this.order)-1] {
			this.order = append(this.order, x)
		}
	}
}

// Pop pop
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	x := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if x == this.order[len(this.order)-1] {
		this.order = this.order[:len(this.order)-1]
	}
}

// Top top
func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

// Min min
func (this *MinStack) Min() int {
	if len(this.order) == 0 {
		return 0
	}
	return this.order[len(this.order)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.Top(), obj.Min()) // -3 -3
	obj.Pop()
	fmt.Println(obj.Top(), obj.Min()) // 0 -2

	fmt.Println()
	obj2 := Constructor()
	obj2.Push(-2)
	obj2.Push(0)
	obj2.Push(-1)
	fmt.Println(obj2.Min(), obj2.Top()) // -2 -1
	obj2.Pop()
	fmt.Println(obj2.Min(), obj2.Top()) // -2 0

	fmt.Println()
	obj3 := Constructor()
	obj3.Push(2)
	obj3.Push(0)
	obj3.Push(3)
	obj3.Push(0)
	fmt.Println(obj3.Min(), obj3.Top()) // 0 0
	obj3.Pop()
	fmt.Println(obj3.Min(), obj3.Top()) // 0 3
	obj3.Pop()
	fmt.Println(obj3.Min(), obj3.Top()) // 0 0
	obj3.Pop()
	fmt.Println(obj3.Min(), obj3.Top()) // 2 2
}
