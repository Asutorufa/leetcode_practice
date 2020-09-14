package main

import "fmt"

/**
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3

示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：
输入：root = [1,2,2,null,3,null,3]
输出：false

限制：
0 <= 节点个数 <= 1000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func create(a []int, k int) *TreeNode {
	if len(a) == 0 || len(a)-1 < k {
		return nil
	}
	root := &TreeNode{
		Val: a[k],
	}
	root.Left = create(a, k*2+1)
	root.Right = create(a, k*2+2)
	return root
}

func prePrint(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val)
	fmt.Print(" ")
	prePrint(root.Left)
	prePrint(root.Right)
}

func levelPrint(root *TreeNode) {
	que := newQueue()
	que.push(root)
	for que.size() != 0 {
		x := que.pop()
		if x != nil {
			fmt.Print(x.Val, " ")
			que.push(x.Left)
			que.push(x.Right)
		}
	}
	fmt.Println()
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
		return &TreeNode{}
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

// 递归 有点...不好理解
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}

func recur(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil || l.Val != r.Val {
		return false
	}
	return recur(l.Left, r.Right) && recur(l.Right, r.Left)
}

func main() {
	prePrint(create([]int{1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 5}, 0))
	fmt.Println()
	levelPrint(create([]int{1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 5}, 0))

	fmt.Println(isSymmetric((create([]int{1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 5}, 0))))
	fmt.Println(isSymmetric((create([]int{1, 2, 2, 3, 4, 4, 3, 5, 7, 7, 8, 8, 7, 6, 5}, 0))))
}
