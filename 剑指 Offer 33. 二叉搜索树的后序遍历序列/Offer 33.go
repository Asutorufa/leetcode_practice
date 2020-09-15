package main

import (
	"fmt"
	"math"
)

/**
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。
如果是则返回 true，否则返回 false。
假设输入的数组的任意两个数字都互不相同。

参考以下这颗二叉搜索树：
     5
    / \
   2   6
  / \
 1   3

示例 1：
输入: [1,6,3,2,5]
输出: false

示例 2：
输入: [1,3,2,6,5]
输出: true

提示：
    数组长度 <= 1000
*/

func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p
	for postorder[p] > postorder[j] {
		p++
	}
	// 此处的 p==j 来判断是否为二叉搜索树
	return p == j && recur(postorder, i, m-1) && recur(postorder, m, j-1)
}

/**
比较难想到
利用栈
通过循环不断寻找根节点 如果有节点比当前根节点还要大 表明这不是一个二叉搜索树 返回false
直接循环结束 说明是一个正常的二叉搜索树 返回true
*/
func verifyPostorder2(postorder []int) bool {
	sta := &stack{}
	root := math.MaxInt64
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		for sta.size() != 0 && sta.peek() > postorder[i] {
			root = sta.pop()
		}
		sta.push(postorder[i])
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
	fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5}))  // false
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))  // true
	fmt.Println(verifyPostorder2([]int{1, 6, 3, 2, 5})) // false
	fmt.Println(verifyPostorder2([]int{1, 3, 2, 6, 5})) // true
}
