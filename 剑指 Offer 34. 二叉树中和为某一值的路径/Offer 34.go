package main

import "fmt"

/**
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。
从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

示例:
给定如下二叉树，以及目标和 sum = 22，
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1

返回:
[
   [5,4,11,2],
   [5,8,4,5]
]

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

func pathSum(root *TreeNode, sum int) [][]int {
	p := &paths{}
	p.recur(root, sum)
	return p.res
}

type paths struct {
	res  [][]int
	path []int
}

/**
* 有点难度
* 利用递归进行不断回溯
* 注意其中一部分要进行深复制, 不然改变path也会改变res中添加过的path
 */
func (p *paths) recur(root *TreeNode, current int) {
	if root == nil {
		return
	}
	p.path = append(p.path, root.Val) // 将当前值加入路径中
	current -= root.Val
	if current == 0 && root.Left == nil && root.Right == nil { // 判断路径是否为 sum(->也就是current减为0时)
		//deepcopy -> 此处需要对path进行深复制, 不可浅复制
		s := make([]int, len(p.path))
		copy(s, p.path)
		p.res = append(p.res, s)
	}
	fmt.Println("current", current, "root.Val", root.Val, "path", p.path, "res", p.res)
	p.recur(root.Left, current) // 递归对左子树和右子树的路径进行判断
	p.recur(root.Right, current)
	p.path = p.path[:len(p.path)-1] // 回溯路径
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

func main() {
	levelPrint(create([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 0, 5, 1}, 0))
	fmt.Println(pathSum(create([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 0, 5, 1}, 0), 22))
	/**
	current 17 root.Val 5 path [5] res []
	current 13 root.Val 4 path [5 4] res []
	current 2 root.Val 11 path [5 4 11] res []
	current -5 root.Val 7 path [5 4 11 7] res []
	current 0 root.Val 2 path [5 4 11 2] res [[5 4 11 2]]
	current 9 root.Val 8 path [5 8] res [[5 4 11 2]]
	current -4 root.Val 13 path [5 8 13] res [[5 4 11 2]]
	current 5 root.Val 4 path [5 8 4] res [[5 4 11 2]]
	current 0 root.Val 5 path [5 8 4 5] res [[5 4 11 2] [5 8 4 5]]
	current 4 root.Val 1 path [5 8 4 1] res [[5 4 11 2] [5 8 4 5]]
	[[5 4 11 2] [5 8 4 5]]
	*/
}
