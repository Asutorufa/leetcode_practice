package main

import "fmt"

/**
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：
输入：matrix = [
		[1,2,3],
		[4,5,6],
		[7,8,9]
	]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [
		[1,2,3,4],
		[5,6,7,8],
		[9,10,11,12]
	]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

限制：
    0 <= matrix.length <= 100
    0 <= matrix[i].length <= 100
*/

// 比较简单... 有点麻烦, 要多考虑考虑, 尤其最后一个例子(-> 为什么要用两个循环)
func spiralOrder(matrix [][]int) []int {
	lx := len(matrix)
	if lx == 0 {
		return []int{}
	}
	ly := len(matrix[0])

	capacity := lx * ly
	count := 0
	x, y := 0, 0

	record := make([][]bool, lx)
	for index := range record {
		record[index] = make([]bool, ly)
	}

	var res []int

	add := func() {
		// fmt.Println(matrix[x][y], x, y, count)
		res = append(res, matrix[x][y])
		record[x][y] = true
		count++
	}

	add()
	for count != capacity {
		for !(x-1 < 0) && !record[x-1][y] {
			x--
			add()
		}
		for !(y+1 >= ly) && !record[x][y+1] {
			y++
			add()
		}
		for !(x+1 >= lx) && !record[x+1][y] {
			x++
			add()
		}
		for !(y-1 < 0) && !record[x][y-1] {
			y--
			add()
		}
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	})) // [1,2,3,4,8,12,11,10,9,5,6,7]
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})) // [1 2 3 6 9 8 7 4 5]
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})) // [1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10]
}
