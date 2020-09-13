package main

import (
	"fmt"
)

/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

限制：
0 <= 节点个数 <= 5000
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func print(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func create(x []int) *ListNode {
	root := &ListNode{}
	tmp := root
	for index := range x {
		tmp.Val = x[index]
		if tmp.Next == nil && index != len(x)-1 {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}
	return root
}

/**
双指针: 很巧妙 想不到
*/
func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	for head != nil {
		tmp := cur
		cur = head
		head = head.Next
		cur.Next = tmp
	}
	return cur
}

/**
my solution:循环一次 新建一个链表
*/
func reverseList2(head *ListNode) (root *ListNode) {
	for head != nil {
		root = &ListNode{
			Val:  head.Val,
			Next: root,
		}
		head = head.Next
	}
	return
}

func main() {
	x := create([]int{1, 2, 3, 4, 5})
	print(x)
	y := reverseList(x)
	print(y)
	x = create([]int{1, 2, 3, 4, 5})
	print(x)
	z := reverseList2(x)
	print(z)
}
