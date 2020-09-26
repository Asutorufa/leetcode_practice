package main

import "fmt"

/**
输入一棵二叉树的根节点，求该树的深度。
从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

提示：
    节点总数 <= 10000
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

// 简单题 层次遍历就能解决... 一次写成功
func maxDepth(root *TreeNode) int {
	return level(root)
}

func level(root *TreeNode) (res int) {
	if root == nil {
		return
	}
	queue := &queue{}
	queue.push(root)
_retry:
	i := queue.size()
	if i == 0 {
		// fmt.Println()
		return
	}
	res++
	// fmt.Println("Level ", res)
	for j := 0; j < i; j++ {
		x := queue.pop()
		// fmt.Println(x.Val)
		if x.Left != nil {
			queue.push(x.Left)
		}
		if x.Right != nil {
			queue.push(x.Right)
		}
	}
	goto _retry
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

func (q *queue) size() int {
	return len(q.qu)
}

func main() {
	fmt.Println(level(create([]int{5, 3, 6, 2, 4, 0, 0, 1}, 0)))
}
