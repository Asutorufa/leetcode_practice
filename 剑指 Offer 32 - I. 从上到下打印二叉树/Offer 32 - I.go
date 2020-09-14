package main

/**
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

例如:
给定二叉树: [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7

返回：
[3,9,20,15,7]

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指Offer 28写过了 就不测试了
func levelOrder(root *TreeNode) (res []int) {
	que := newQueue()
	que.push(root)
	for que.size() != 0 {
		x := que.pop()
		if x != nil {
			res = append(res, x.Val)
			que.push(x.Left)
			que.push(x.Right)
		}
	}
	return res
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

func main() {}
