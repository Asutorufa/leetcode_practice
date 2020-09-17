package main

import (
	"bytes"
	"fmt"
)

/**
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

示例:
输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

限制：
1 <= s 的长度 <= 8
*/

func permutation(s string) []string {
	c := &change{
		str: s,
		res: []string{},
	}
	c.dfs(0)
	return c.res
}

type change struct {
	str string
	res []string
}

func (c *change) dfs(x int) {
	if x == len(c.str)-1 { // 检测当前递归长度是否与字符串相同
		c.res = append(c.res, c.str) // 将递归中的字符串加入结果集中
		return
	}
	r := []byte{}
	for i := x; i < len(c.str); i++ {
		if bytes.Contains(r, []byte{c.str[i]}) { // 检测是否已经包含当前字符串
			continue
		}
		r = append(r, c.str[i]) // 将当前字符串加入已递归集中
		// fmt.Println(string(r), c.str)
		// fmt.Print("before ", c.str, " change ", x, i)
		c.swap(x, i) // 交换 x 和 i 位置的字符
		// fmt.Print(" after ", c.str, "\n")
		c.dfs(x + 1) // 递归
		c.swap(i, x) // 撤销交换
	}
}

func (c *change) swap(x, i int) {
	if x == i {
		return
	}
	z := []byte(c.str)
	tmp := z[i]
	z[i] = z[x]
	z[x] = tmp
	c.str = string(z)
}

func main() {
	fmt.Println(permutation("abc"))
}
