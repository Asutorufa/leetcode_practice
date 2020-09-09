package main

import (
	"fmt"
)

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，
它每次可以向左、右、上、下移动一格（不能移动到方格外），
也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。
但它不能进入方格 [35, 38]，因为3+5+3+8=19。
请问该机器人能够到达多少个格子？

示例 1：

输入：m = 2, n = 3, k = 1
输出：3

示例 2：

输入：m = 3, n = 1, k = 0
输出：1

提示：

    1 <= n,m <= 100
    0 <= k <= 20
*/

func s(n int) (sum int) {
	for n/10 != 0 || n%10 != 0 {
		sum += n % 10
		n /= 10
	}
	return
}

/*
	dfs
*/

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for x := range visited {
		visited[x] = make([]bool, n)
	}
	return dfs(0, 0, 0, 0, m, n, k, visited)
}

func dfs(i, j, si, sj, m, n, k int, visited [][]bool) int {
	if i >= m || j >= n || k < si+sj || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	return 1 + dfs(i+1, j, s(i+1), sj, m, n, k, visited) + dfs(i, j+1, si, s(j+1), m, n, k, visited)
}

/*
	bfs <- 将dfs中的递归换为队列
*/

func movingCountBfs(m int, n int, k int) int {
	var res int
	queue := newQueue()
	queue.push([4]int{0, 0, 0, 0})

	visited := make([][]bool, m)
	for x := range visited {
		visited[x] = make([]bool, n)
	}

	for queue.size() > 0 {
		x := queue.pop()
		i := x[0]
		j := x[1]
		si := x[2]
		sj := x[3]
		if i >= m || j >= n || k < si+sj || visited[i][j] {
			continue
		}
		visited[i][j] = true
		res++
		queue.push([4]int{i + 1, j, s(i + 1), sj})
		queue.push([4]int{i, j + 1, si, s(j + 1)})
	}
	return res
}

type queue struct {
	qu [][4]int
}

func newQueue() *queue {
	return &queue{
		qu: [][4]int{},
	}
}

func (q *queue) pop() [4]int {
	if len(q.qu) == 0 {
		return [4]int{}
	}
	x := q.qu[0]
	q.qu = q.qu[1:]
	return x
}

func (q *queue) push(x [4]int) {
	q.qu = append(q.qu, x)
}

func (q *queue) size() int {
	return len(q.qu)
}

func main() {
	fmt.Println("dfs")
	fmt.Println(movingCount(2, 3, 1))
	fmt.Println(movingCount(3, 1, 0))
	fmt.Println(movingCount(16, 8, 4))
	fmt.Println(movingCount(38, 15, 9))
	fmt.Println("bfs")
	fmt.Println(movingCountBfs(2, 3, 1))
	fmt.Println(movingCountBfs(3, 1, 0))
	fmt.Println(movingCountBfs(16, 8, 4))
	fmt.Println(movingCountBfs(38, 15, 9))
	fmt.Println("queue test")
	x := newQueue()
	x.push([4]int{1, 1, 1, 1})
	x.push([4]int{2, 2, 2, 2})
	fmt.Println(x.size())
	fmt.Println(x.pop())
	fmt.Println(x.size())
	fmt.Println(x.pop())
	fmt.Println(x.size())
}
