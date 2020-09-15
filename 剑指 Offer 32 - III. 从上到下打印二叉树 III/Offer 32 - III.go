package main

import "fmt"

/**
请实现一个函数按照之字形顺序打印二叉树，
即第一行按照从左到右的顺序打印，
第二层按照从右到左的顺序打印，
第三行再按照从左到右的顺序打印，
其他行以此类推。

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
  [20,9],
  [15,7]
]

提示：
    节点总数 <= 1000
*/

// TreeNode tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
比较简单... 跟第二个差不多 中间加个判断奇偶层就行了...
*/
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	que := newQueue2()
	que.push(root)
	count := 1
	for que.size() != 0 {
		var tmp []int
		// fmt.Println("queue size", que.size())
		for i := que.size(); i > 0; i-- {
			// 这里加个循环 (注意这里不是用队列是否为空来循环, 因为会不断加入子元素到队列中)
			// 会把当前队列中(或者可以说是当前这一层的)所有元素的子元素全部加入队列中
			// 相当巧妙
			x := que.pop()
			if x != nil {
				if count%2 != 0 {
					tmp = append(tmp, x.Val)
				} else if count%2 == 0 {
					tmp = append([]int{x.Val}, tmp...)
				}
				if x.Left != nil {
					que.push(x.Left)
				}
				if x.Right != nil {
					que.push(x.Right)
				}
			}
		}
		res = append(res, tmp)
		count++
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
}
