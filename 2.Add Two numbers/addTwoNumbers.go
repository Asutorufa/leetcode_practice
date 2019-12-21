package main

import "log"

import "container/list"

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	last := root
	for tmp := 0; ; {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := tmp + a + b
		tmp = sum / 10
		last.Val = sum % 10
		if l1 == nil && l2 == nil && tmp == 0 {
			break
		} else {
			last.Next = &ListNode{}
			last = last.Next
		}
	}
	return root
}

func two() {
	one := list.New()
	one.PushBack(1)
	one.PushBack(8)
	// one.PushBack(3)
	two := list.New()
	two.PushBack(0)
	// two.PushBack(5)
	// two.PushBack(6)
	// two.PushBack(4)
	last := list.New()
	for e, f, tmp := one.Front(), two.Front(), 0; ; {
		a, b := 0, 0
		if e != nil {
			a = e.Value.(int)
			e = e.Next()
		}
		if f != nil {
			b = f.Value.(int)
			f = f.Next()
		}
		sum := tmp + a + b
		tmp = sum / 10
		last.PushBack(sum % 10)
		log.Println(a, b)
		if e == nil && f == nil && tmp == 0 {
			break
		}
	}
	for e := last.Front(); e != nil; e = e.Next() {
		log.Println(e.Value)
	}
}

func main() {
	two()
}
