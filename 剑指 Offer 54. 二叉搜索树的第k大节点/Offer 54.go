package main

import (
	"fmt"
)

/**
给定一棵二叉搜索树，请找出其中第k大的节点。

示例 1:
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4

示例 2:
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 简单题 魔改中序遍历 一次性写出来了
func kthLargest(root *TreeNode, k int) int {
	val := new(int)
	search(root,k,val,new(int))
	return *val
}

func search(r *TreeNode,k int,res,now *int){
	if r == nil {
		return
	}
	search(r.Right,k,res,now)
	*now++
	if *now > k{
		return
	}
	if *now == k{
		*res =r.Val
		// fmt.Println("find ",k,r.Val)
		return
	}
	// fmt.Println(r.Val)
	search(r.Left,k,res,now)
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

func main(){
	a := new(int)
	search(create([]int{5,3,6,2,4,0,0,1},0),4,a,new(int))
	fmt.Println(*a)
	fmt.Println()
	fmt.Println(kthLargest(create([]int{5,3,6,2,4,0,0,1},0),3))
	fmt.Println()
	fmt.Println(kthLargest(create([]int{3,1,4,0,2},0),1))
}
