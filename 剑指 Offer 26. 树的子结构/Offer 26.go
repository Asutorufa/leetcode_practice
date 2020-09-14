package main

import "fmt"

/**
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:
     3
    / \
   4   5
  / \
 1   2

给定的树 B：
   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：
输入：A = [1,2,3], B = [3,1]
输出：false

示例 2：
输入：A = [3,4,5,1,2], B = [4,1]
输出：true

限制：
0 <= 节点个数 <= 10000
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
这题也比较简单, 就是麻烦点... 大量运用递归

可以不用getTree 来寻找 直接使用(isSubStructure)递归A.Left 和 A.Right来不断判断
这里getTree是我第一次自己写时用的, 后来看了解法, 发现还可以更简单
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	x := getTree(A, B.Val)
	if checkTree(x, B) {
		return true
	}
	if x != nil && isSubStructure(x.Left, B) {
		return true
	}
	if x != nil && isSubStructure(x.Right, B) {
		return true
	}
	return false
}

func checkTree(A *TreeNode, B *TreeNode) bool {
	if A == nil && B != nil {
		return false
	}
	if B == nil {
		return true
	}
	if A.Val != B.Val {
		return false
	}
	if checkTree(A.Left, B.Left) == false {
		return false
	}
	if checkTree(A.Right, B.Right) == false {
		return false
	}
	return true
}

func getTree(A *TreeNode, Val int) *TreeNode {
	if A == nil {
		return nil
	}
	if A.Val == Val {
		return A
	}
	if l := getTree(A.Left, Val); l != nil {
		return l
	}
	if r := getTree(A.Right, Val); r != nil {
		return r
	}
	return nil
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

func main() {
	x := create([]int{3, 4, 5, 1, 2}, 0)
	// prePrint(x) // 3,4,1,2,5
	fmt.Println(isSubStructure(x, create([]int{4, 1}, 0)))
	y := create([]int{1, 2, 3}, 0)
	fmt.Println(isSubStructure(y, create([]int{3, 1}, 0)))
	z := create([]int{-1, 3, 2, 0}, 0)
	fmt.Println(isSubStructure(z, create([]int{}, 0)))
	c := create([]int{1, 0, 1, -4, -3}, 0)
	fmt.Println(isSubStructure(c, create([]int{1, -4}, 0)))
	v := create([]int{4, 2, 3, 4, 5, 6, 7, 8, 9}, 0)
	fmt.Println(isSubStructure(v, create([]int{4, 8, 9}, 0)))
}
