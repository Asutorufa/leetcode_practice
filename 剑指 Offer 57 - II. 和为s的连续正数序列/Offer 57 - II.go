package main

import "fmt"

/**
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：
输入：target = 9
输出：[[2,3,4],[4,5]]

示例 2：
输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

限制：
    1 <= target <= 10^5
*/

// 滑动窗口 窗口只想右滑动
// 大于target就左边界减少 小于target就扩大右边界
// 从题目中的例子可以知道 左边界最大不会超过target/2 所以左边界到target/2就行了
func findContinuousSequence(target int) (res [][]int) {
	i, j := 1, 0 // 因为是正整数, 所以应该初始化为 1
	// 区别1：right初值为0，不是 不取right版本中 的1，也不是2。
	//        因为代码中，单纯是通过 sum +=right来增加right，所以必须保证right经过1
	//        而取right的代码中，是先right++,sum +=right , 因此right初值只能为0
	sum := 0
	for i <= target/2 {
		if sum < target {
			j++
			sum += j
		} else if sum > target {
			sum -= i
			i++
		} else {
			tmp := []int{}
			for x := i; x <= j; x++ {
				tmp = append(tmp, x)
			}
			res = append(res, tmp)
			sum -= i
			i++
		}
	}
	return
}

func main() {
	fmt.Println(findContinuousSequence(9))
	fmt.Println(findContinuousSequence(15))
}
