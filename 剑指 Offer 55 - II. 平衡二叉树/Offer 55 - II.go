package main

import (
	"fmt"
)

/**
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

示例 1:
给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:
给定二叉树 [1,2,2,3,3,null,null,4,4]
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

限制：
    1 <= 树的结点个数 <= 10000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 理解后写 还是比较简单的... 递归左右子树就行了
func isBalanced(root *TreeNode) bool {
	_, is := check(0, root)
	return is
}

func check(level int, root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	x, is := check(level, root.Left)
	if !is {
		return x, is
	}
	y, is := check(level, root.Right)
	if !is {
		return y, is
	}
	if x > y {
		level = x
	} else {
		level = y
	}
	// fmt.Println(level)
	if x-y > 1 || x-y < -1 {
		return level + 1, false
	}
	return level + 1, true
}

func create(a []int, k int) *TreeNode {
	if len(a) == 0 || len(a)-1 < k {
		return nil
	}
	if a[k] == 0 { // 此处可不此操作
		return nil
	}
	root := &TreeNode{
		Val: a[k],
	}
	root.Left = create(a, k*2+1)
	root.Right = create(a, k*2+2)
	return root
}

func main() {
	fmt.Println(isBalanced(create([]int{3, 9, 20, 0, 0, 15, 7}, 0)))                          // true
	fmt.Println(isBalanced(create([]int{1, 2, 2, 3, 3, 0, 0, 4, 4}, 0)))                      // false
	fmt.Println(isBalanced(create([]int{1, 2, 3, 4, 5, 6, 7}, 0)))                            // true
	fmt.Println(isBalanced(create([]int{1, 2, 3, 4, 5, 6, 7, 8}, 0)))                         // true
	fmt.Println(isBalanced(create([]int{1, 2, 3, 4, 5, 6, 7, 8, 0, 0, 0, 0, 0, 0, 0, 9}, 0))) // false
	fmt.Println(isBalanced(create([]int{1, 0, 2, 0, 0, 0, 3}, 0)))                            // false
	fmt.Println(isBalanced(create([]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 0, 0, 5, 5}, 0)))    // true
}

/**
错误理解
*/

// 错误理解平衡二叉树
// 这个是不完全二叉树
func isBalanced2(root *TreeNode) bool {
	queue := &queue{}
	queue.push(root)
	count2 := 0
	not := false
_cont:
	not = false
	s := queue.size()
	if s == 0 {
		goto _check
	}

	for i := 0; i < s; i++ {
		x := queue.pop()
		if x == nil {
			not = true
			continue
		}
		// fmt.Println(x.Val)
		queue.push(x.Left)
		queue.push(x.Right)
	}
	if not {
		count2++
	}
	// fmt.Println("count", count2)
	goto _cont
_check:
	if count2 > 2 {
		return false
	}
	return true
}

type queue struct {
	qu []*TreeNode
}

func newQueue() *queue {
	return &queue{
		qu: []*TreeNode{},
	}
}

func (q *queue) pop() *TreeNode {
	if len(q.qu) == 0 {
		return nil
	}
	x := q.qu[0]
	q.qu = q.qu[1:]
	return x
}

func (q *queue) push(x *TreeNode) {
	q.qu = append(q.qu, x)
}

func (q *queue) peek() *TreeNode {
	if len(q.qu) == 0 {
		return nil
	}
	return q.qu[0]
}

func (q *queue) size() int {
	return len(q.qu)
}
