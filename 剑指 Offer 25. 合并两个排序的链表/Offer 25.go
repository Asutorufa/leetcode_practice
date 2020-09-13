package main

import (
	"fmt"
)

/**
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

限制：
0 <= 链表长度 <= 1000
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
	if len(x) == 0 {
		return nil
	}
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
比较简单 一次就写成了
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l *ListNode = &ListNode{}
	root := l
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l.Next = l2
			break
		}
		if l2 == nil {
			l.Next = l1
			break
		}
		for l2 != nil && l1.Val >= l2.Val {
			l.Next = l2
			l = l.Next
			l2 = l2.Next
		}
		l.Next = l1
		l = l.Next
		l1 = l1.Next
	}
	return root.Next
}

func main() {
	x := mergeTwoLists(create([]int{1, 2, 4}), create([]int{1, 3, 4}))
	print(x)
	x = mergeTwoLists(create([]int{1, 5}), create([]int{1, 2, 3, 4, 5}))
	print(x)
	x = mergeTwoLists(create([]int{}), create([]int{0}))
	print(x)
}
