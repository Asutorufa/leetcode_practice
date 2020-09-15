package main

import "fmt"

/**
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

例如:
给定二叉树: [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7

返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]

提示：
    节点总数 <= 1000
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

/**
比较简单... 修改一下32-I存储的队列就行了
*/
func levelOrder(root *TreeNode) (res [][]int) {
	que := newQueue()
	que.push(root, 0)
	for que.size() != 0 {
		x, level := que.pop()
		if x != nil {
			// fmt.Println(x.Val, level)
			if level > len(res)-1 {
				res = append(res, []int{x.Val})
			} else {
				res[level] = append(res[level], x.Val)
			}
			que.push(x.Left, level+1)
			que.push(x.Right, level+1)
		}
	}
	return res
}

type qu struct {
	node  *TreeNode
	level int
}
type queue struct {
	qu []*qu
}

func newQueue() *queue {
	return &queue{
		qu: []*qu{},
	}
}

func (q *queue) pop() (*TreeNode, int) {
	if len(q.qu) == 0 {
		return &TreeNode{}, 0
	}
	x := q.qu[0]
	q.qu = q.qu[1:]
	return x.node, x.level
}

func (q *queue) push(x *TreeNode, level int) {
	q.qu = append(q.qu, &qu{
		node:  x,
		level: level,
	})
}

func (q *queue) size() int {
	return len(q.qu)
}

/**
利用队列的大小来判断层次... 比较难理解
*/
func levelOrder2(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	que := newQueue2()
	que.push(root)
	for que.size() != 0 {
		var tmp []int
		// fmt.Println("queue size", que.size())
		for i := que.size(); i > 0; i-- {
			// 这里加个循环 (注意这里不是用队列是否为空来循环, 因为会不断加入子元素到队列中)
			// 会把当前队列中(或者可以说是当前这一层的)所有元素的子元素全部加入队列中
			// 相当巧妙
			x := que.pop()
			if x != nil {
				tmp = append(tmp, x.Val)
				if x.Left != nil {
					que.push(x.Left)
				}
				if x.Right != nil {
					que.push(x.Right)
				}
			}
		}
		res = append(res, tmp)
	}
	return res
}

type queue2 struct {
	qu []*TreeNode
}

func newQueue2() *queue2 {
	return &queue2{
		qu: []*TreeNode{},
	}
}

func (q *queue2) pop() *TreeNode {
	if len(q.qu) == 0 {
		return &TreeNode{}
	}
	x := q.qu[0]
	q.qu = q.qu[1:]
	return x
}

func (q *queue2) push(x *TreeNode) {
	q.qu = append(q.qu, x)
}

func (q *queue2) size() int {
	return len(q.qu)
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

func main() {
	fmt.Println(levelOrder(create([]int{1, 2, 3, 4, 5}, 0)))
	fmt.Println(levelOrder(create([]int{3, 9, 20, 15, 17}, 0)))
	fmt.Println(levelOrder2(create([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)))
	fmt.Println(levelOrder2(create([]int{3, 9, 20, 15, 17}, 0)))
}
