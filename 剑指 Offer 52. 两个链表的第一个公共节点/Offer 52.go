package main

import (
	"fmt"
	"math"
)

/**
输入两个链表，找出它们的第一个公共节点。

如下面的两个链表：
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_statement.png)
在节点 c1 开始相交。

示例 1：
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_example_1.png)
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。
		从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
		在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

示例 2：
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_example_2.png)
输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Reference of the node with value = 2
输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。
		从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，
		相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

示例 3：
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_example_3.png)
输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
		由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
解释：这两个链表不相交，因此返回 null。

注意：
    如果两个链表没有交点，返回 null.
    在返回结果后，两个链表仍须保持原有的结构。
    可假定整个链表结构中没有循环。
    程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
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

// 很神奇的方法
/**
我们使用两个指针 node1，node2 分别指向两个链表 headA，headB 的头结点，
然后同时分别逐结点遍历，当 node1 到达链表 headA 的末尾时，重新定位到链表 headB 的头结点；
当 node2 到达链表 headB 的末尾时，重新定位到链表 headA 的头结点。

这样，当它们相遇时，所指向的结点就是第一个公共结点。
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

// 自己写的 巨慢
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB

	al := getLen(a)
	bl := getLen(b)

	if al-bl > 0 {
		x := math.Abs(float64(al - bl))
		for i := float64(0); i < x; i++ {
			a = a.Next
		}
	} else {
		x := math.Abs(float64(al - bl))
		for i := float64(0); i < x; i++ {
			b = b.Next
		}

	}
	for a != nil && b != nil {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil
}

func getLen(a *ListNode) (res int) {
	for a != nil {
		res++
		a = a.Next
	}
	return
}

func create(a []int, b []int, common []int) (*ListNode, *ListNode) {
	roota := &ListNode{}
	rootb := &ListNode{}
	tmpa := roota
	tmpb := rootb

	for index := range a {
		if index != 0 {
			tmpa.Next = &ListNode{}
			tmpa = tmpa.Next
		}
		tmpa.Val = a[index]
	}

	for index := range b {
		if index != 0 {
			tmpb.Next = &ListNode{}
			tmpb = tmpb.Next
		}
		tmpb.Val = b[index]
	}

	rootc := &ListNode{}
	tmpa.Next = rootc
	tmpb.Next = rootc

	for index := range common {
		if index != 0 {
			rootc.Next = &ListNode{}
			rootc = rootc.Next
		}
		rootc.Val = common[index]
	}

	return roota, rootb
}

func print(a *ListNode) {
	for a != nil {
		fmt.Print(a.Val, " ")
		a = a.Next
	}
	fmt.Println()
}

func main() {
	print(getIntersectionNode(create([]int{4, 1}, []int{5, 0, 1}, []int{8, 4, 5})))
	print(getIntersectionNode2(create([]int{4, 1}, []int{5, 0, 1}, []int{8, 4, 5})))
}
