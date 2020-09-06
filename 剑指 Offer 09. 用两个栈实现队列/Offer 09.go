package main

import (
	"log"
)

/*
用两个栈实现一个队列。
队列的声明如下，
请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]

示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]

提示：

    1 <= values <= 10000
    最多会对 appendTail、deleteHead 进行 10000 次调用
*/

type CQueue struct {
	stack1, stack2 *stack
}

func Constructor() CQueue {
	return CQueue{
		stack1: newStack(),
		stack2: newStack(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.isEmpty() { // only if the stack2 is empty, trans stack1 to stack2
		for {
			if this.stack1.isEmpty() {
				break
			}
			this.stack2.push(this.stack1.pop())
		}
	}
	if this.stack2.isEmpty() {
		return -1
	}
	return this.stack2.pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

// stack defines Stack1 stack
type stack struct {
	nums []int
}

//  newStack creates a empty stack
func newStack() *stack {
	return &stack{
		nums: []int{},
	}
}

func (s *stack) push(n int) {
	s.nums = append(s.nums, n)
}

func (s *stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

func (s *stack) len() int {
	return len(s.nums)
}

func (s *stack) isEmpty() bool {
	return s.len() == 0
}

func main() {
	obj := Constructor()
	log.Println(obj.DeleteHead())
	obj.AppendTail(2)
	obj.AppendTail(3)
	log.Println(obj.DeleteHead())
	log.Println(obj.DeleteHead())
	log.Println(obj.DeleteHead())
	obj.AppendTail(3)
	log.Println(obj.DeleteHead())
	log.Println(obj.DeleteHead())
	log.Println(obj.DeleteHead())
}
