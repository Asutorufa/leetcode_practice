package main

import (
	"fmt"
)

/*
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

提示：

    1 <= board.length <= 200
    1 <= board[i].length <= 200
*/

func exist(board [][]byte, word string) bool {
	for x := range board {
		for y := range board[x] {
			if dfs(board, word, x, y, 0) {
				return true
			}
		}
	}
	return false
}

/*
算法原理：
    深度优先搜索： 可以理解为暴力法遍历矩阵中所有字符串可能性。DFS 通过递归，先朝一个方向搜到底，再回溯至上个节点，沿另一个方向搜索，以此类推。
    剪枝： 在搜索中，遇到 这条路不可能和目标字符串匹配成功 的情况（例如：此矩阵元素和目标字符不同、此元素已被访问），则应立即返回，称之为 可行性剪枝 。

算法剖析：
    递归参数： 当前元素在矩阵 board 中的行列索引 i 和 j ，当前目标字符在 word 中的索引 k 。
    终止条件：
        返回 false ： ① 行或列索引越界 或 ② 当前矩阵元素与目标字符不同 或 ③ 当前矩阵元素已访问过 （③ 可合并至 ② ） 。
        返回 true ： 字符串 word 已全部匹配，即 k = len(word) - 1 。
    递推工作：
        标记当前矩阵元素： 将 board[i][j] 值暂存于变量 tmp ，并修改为字符 '/' ，代表此元素已访问过，防止之后搜索时重复访问。
        搜索下一单元格： 朝当前元素的 上、下、左、右 四个方向开启下层递归，使用 或 连接 （代表只需一条可行路径） ，并记录结果至 res 。
        还原当前矩阵元素： 将 tmp 暂存值还原至 board[i][j] 元素。
    回溯返回值： 返回 res ，代表是否搜索到目标字符串。
*/

func dfs(board [][]byte, word string, i, j, k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	tmp := board[i][j]
	board[i][j] = '/' // 此处使用 / 进行占位, 但有可能字符串中就有 /, 所以应使用一个bool二维数组存储访问过的路径
	res := dfs(board, word, i+1, j, k+1) ||
		dfs(board, word, i-1, j, k+1) ||
		dfs(board, word, i, j+1, k+1) ||
		dfs(board, word, i, j-1, k+1)
	board[i][j] = tmp
	return res
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
	board = [][]byte{
		{'a', 'd'},
		{'b', 'c'},
	}
	word = "abcd"
	fmt.Println(exist(board, word))
	board = [][]byte{{'a'}}
	word = "a"
	fmt.Println(exist(board, word))
	board = [][]byte{{'a', 'a'}}
	word = "aaa"
	fmt.Println(exist(board, word))
	board = [][]byte{{'a', 'b'}}
	word = "ba"
	fmt.Println(exist(board, word))
	board = [][]byte{{'a', 'b'}, {'c', 'd'}}
	word = "cdba"
	fmt.Println(exist(board, word))
	board = [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}
	word = "AAB"
	fmt.Println(exist(board, word))

	board = [][]byte{{'a'}}
	word = "ab"
	fmt.Println(exist(board, word))

}
