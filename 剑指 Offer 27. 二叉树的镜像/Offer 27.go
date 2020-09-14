package main

import "fmt"

/**
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1

示例 1：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

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

/**
比较简单... 不多说了
*/
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)
	root.Right = left
	root.Left = right
	return root
}

func main() {
	prePrint(mirrorTree(create([]int{4, 2, 7, 1, 3, 6, 9}, 0))) // [4,7,2,9,6,3,1] -> 4 7 9 6 2 3 1
	fmt.Println()
}
