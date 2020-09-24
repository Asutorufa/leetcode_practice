package main

import "fmt"

/**
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

示例 1:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

提示：
    0 < grid.length <= 200
    0 < grid[0].length <= 200
*/

/**
动态规划
设 f(i,j)为从棋盘左上角走至单元格 (i,j) 的礼物最大累计价值，易得到以下递推关系：
	f(i,j) 等于 f(i,j−1) 和 f(i−1,j) 中的较大值加上当前单元格礼物价值 grid(i,j) 。

f(i,j)=max⁡[f(i,j−1),f(i−1,j)]+grid(i,j)
*/
func maxValue(grid [][]int) int {
	w := len(grid)
	h := len(grid[0])
	for i := 1; i < w; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < h; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < w; i++ {
		for j := 1; j < h; j++ {
			if grid[i][j-1] > grid[i-1][j] {
				grid[i][j] += grid[i][j-1]
			} else {
				grid[i][j] += grid[i-1][j]
			}
		}
	}
	return grid[w-1][h-1]
}

// 自己写的, 超出时间限制.....
func maxValue2(grid [][]int) int {
	w := len(grid)
	if w == 0 {
		return 0
	}
	h := len(grid[0])
	return dfs(0, 0, w, h, 0, grid)
}

// 超出时间限制.....
func dfs(i, j int, w, h int, now int, grid [][]int) int {
	if i >= w || j >= h {
		return now
	}
	// fmt.Println(i, j, grid[i][j])
	n := grid[i][j]
	now += n
	x := dfs(i+1, j, w, h, now, grid)
	y := dfs(i, j+1, w, h, now, grid)
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxValue([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))

	fmt.Println(maxValue2([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}
