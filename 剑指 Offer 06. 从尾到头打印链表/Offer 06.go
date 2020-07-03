package main

import "log"

/*
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]



限制：

0 <= 链表长度 <= 10000
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
// 反过来就行了 类似栈
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	headTmp := head
	pLeft := &ListNode{}
	pRight := &ListNode{}
	pRight.Val = headTmp.Val
	for headTmp.Next != nil {
		pLeft = pRight
		pRight = &ListNode{}
		pRight.Val = headTmp.Next.Val
		pRight.Next = pLeft
		headTmp = headTmp.Next
	}
	//pRight = pRight.Next
	var ret []int
	for pRight != nil {
		//log.Println(pRight.Val)
		ret = append(ret, pRight.Val)
		pRight = pRight.Next
	}
	return ret
}

func main() {
	test := &ListNode{}
	head := test
	for _, x := range []int{1, 2, 3, 5, 6, 7, 8, 9} {
		test.Next = &ListNode{}
		test = test.Next
		test.Val = x
	}
	head = head.Next
	log.Println(reversePrint(head))
}
