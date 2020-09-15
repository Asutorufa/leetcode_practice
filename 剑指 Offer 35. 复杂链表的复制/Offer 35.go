package main

/**
请实现 copyRandomList 函数，复制一个复杂链表。
在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，
还有一个 random 指针指向链表中的任意节点或者 null。

示例 1：
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e1.png)
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

示例 2：
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e2.png)
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]

示例 3：
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e3.png)
输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]

示例 4：
输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 null。

提示：
    -10000 <= Node.val <= 10000
    Node.random 为空（null）或指向链表中的节点。
    节点数目不超过 1000 。
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
相当巧妙的方法 看了就懂了
*/
func copyRandomList(head *Node) *Node {
	s := map[*Node]*Node{}
	cur := head
	for cur != nil {
		s[cur] = &Node{
			Val: cur.Val,
		}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		s[cur].Next = s[cur.Next]
		s[cur].Random = s[cur.Random]
		cur = cur.Next
	}
	return s[head]
}

/**
[](https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/solution/fu-za-lian-biao-de-fu-zhi-jian-dan-yi-dong-de-san-/)
先将当前的每个节点复制一遍 复制到当前节点的后面
	比如: 1 -> 2 -> 3 复制完: 1 -> 1 -> 2 -> 2 -> 3 -> 3
然后复制random:
	相当于是 当前节点的random所指的节点的下一个就是新的random所指的节点 <- 根据上面的复制
最后把旧节点和新节点进行拆分就行了
*/
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil { // 复制
		tmp := &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next = tmp
		cur = tmp.Next
	}

	cur = head
	for cur != nil { // 复制 random
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	cur = head
	res := cur.Next
	for cur != nil { // 拆分
		clone := cur.Next
		cur.Next = clone.Next
		if clone.Next != nil {
			clone.Next = clone.Next.Next
		}
		cur = cur.Next
	}
	return res
}

func main() {}
