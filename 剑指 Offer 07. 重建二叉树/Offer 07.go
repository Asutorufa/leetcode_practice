package main

import (
	"fmt"
)

/*
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]

返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

限制：

0 <= 节点个数 <= 5000
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 { // check array's length
		return nil
	}
	root := &TreeNode{ // create root
		Val: preorder[0],
	}

	var left int
	for left = range inorder {
		if inorder[left] == root.Val {
			break
		}
	} // get root's left tree length

	right := len(preorder) - left - 1 // get root's right tree length
	// fmt.Println("left:", left, "right:", right)
	if left > 0 {
		// fmt.Println("left child", preorder[1:left+1], inorder[:left])
		root.Left = buildTree(preorder[1:left+1], inorder[:left]) // use recursion
		// preorder[1:left+1] -> left tree of preorder
		// inorder[:left] -> left tree of inorder
	}
	if right > 0 {
		// fmt.Println("right child", preorder[left+1:], inorder[left+1:])
		root.Right = buildTree(preorder[left+1:], inorder[left+1:]) // use recursion
		// preorder[left+1:] -> right tree of preorder
		// inorder[left+1:] -> right tree of preorder
	}
	return root
}

func prePrint(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	prePrint(root.Left)
	prePrint(root.Right)
}

func inPrint(root *TreeNode) {
	if root == nil {
		return
	}
	inPrint(root.Left)
	fmt.Println(root.Val)
	inPrint(root.Right)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)
	/*
		left: 1 right: 3
		left child [9] [9]
		left: 0 right: 0
		right child [20 15 7] [15 20 7]
		left: 1 right: 1
		left child [15] [15]
		left: 0 right: 0
		right child [7] [7]
		left: 0 right: 0
	*/
	prePrint(tree)
	fmt.Println("------------")
	inPrint(tree)
}
