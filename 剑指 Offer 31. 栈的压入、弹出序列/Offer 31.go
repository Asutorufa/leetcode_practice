package main

import "fmt"

/**
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
假设压入栈的所有数字均不相等。
例如，序列 {1,2,3,4,5} 是某栈的压栈序列，
序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，
但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

示例 1：
输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

示例 2：
输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。

提示：
    0 <= pushed.length == popped.length <= 1000
    0 <= pushed[i], popped[i] < 1000
    pushed 是 popped 的排列。
*/

// 比较简单...一次就成功了 直接用栈就行.. 也可用数组模拟...
func validateStackSequences(pushed []int, popped []int) bool {
	stack := &stack{}
	index := 0
	for _, x := range popped {
		if stack.size() == 0 {
			stack.push(pushed[index])
			index++
		}
	_retry:
		// fmt.Println(stack.s)
		if stack.peek() == x {
			stack.pop()
			continue
		} else {
			if index > len(pushed)-1 {
				return false
			}
			stack.push(pushed[index])
			index++
			goto _retry
		}
	}
	return true
}

type stack struct {
	s []int
}

func (s *stack) push(a int) {
	s.s = append(s.s, a)
}

func (s *stack) pop() int {
	a := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return a
}

func (s *stack) peek() int {
	return s.s[len(s.s)-1]
}

func (s *stack) size() int {
	return len(s.s)
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})) // true
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2})) // false
}
