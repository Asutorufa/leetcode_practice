package main

import "fmt"

/**
输入一个链表，输出该链表中倒数第k个节点。
为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。
这个链表的倒数第3个节点是值为4的节点。

示例：
给定一个链表: 1->2->3->4->5, 和 k = 2.
返回链表 4->5.
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

/**
双指针: 相当巧妙的解法, 想不到
第一个指针先走k步, 然后两个指针同时走, 当第一个指针走到链表未是返回第二个就行了
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var f *ListNode = head
	var l *ListNode = head
	for i := 1; i <= k; i++ {
		f = f.Next
	}
	for f != nil {
		f = f.Next
		l = l.Next
	}
	return l
}

/**
my solution: 用一个数组存每个节点, 返回长度-k的节点, 占点内存
*/
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	var arr []*ListNode
	var count int = 0
	for head != nil {
		arr = append(arr, head)
		count++
		head = head.Next
	}
	return arr[count-k]
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

func main() {
	x := create([]int{1, 2, 3, 4, 5})
	y := getKthFromEnd(x, 2)
	z := getKthFromEnd2(x, 2)
	print(x)
	print(y)
	print(z)
}
