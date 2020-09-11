package main

import "fmt"

/**
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
返回删除后的链表的头节点。

注意：此题对比原题有改动

示例 1:
输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.

示例 2:
输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

说明：
    题目保证链表中节点的值互不相同
    若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 这题比较简单 就不多说了
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	tmp := head
	var last *ListNode = nil
	for head != nil {
		if head.Val == val {
			last.Next = head.Next
			break
		}
		last = head
		head = head.Next
	}
	return tmp
}

func print(head *ListNode) {
	for head.Next != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println()
}

func create(x []int) *ListNode {
	root := &ListNode{}
	tmp := root
	for index := range x {
		tmp.Val = x[index]
		if tmp.Next == nil {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}
	return root
}

func main() {
	x := create([]int{2, 3, 4, 5})
	print(x)
	deleteNode(x, 4)
	print(x)
	y := create([]int{4, 5, 1, 9})
	print(y)
	deleteNode(y, 1)
	print(y)
	y = create([]int{-3, 5, -99})
	print(y)
	y = deleteNode(y, -3)
	print(y)
}
