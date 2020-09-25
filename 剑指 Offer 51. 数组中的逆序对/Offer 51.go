package main

import "fmt"

/**
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
输入一个数组，求出这个数组中的逆序对的总数。

示例 1:
输入: [7,5,6,4]
输出: 5

限制：
0 <= 数组长度 <= 50000
*/

// 暴力 超出时间限制
func reversePairs2(nums []int) int {
	l := len(nums)
	res := 0
	for index := range nums {
		for i := index; i < l; i++ {
			if nums[index] > nums[i] {
				res++
			}
		}
	}
	return res
}

func sort(arr []int) {
	l := len(arr)
	tmp := make([]int, l)
	sortt(arr, 0, l-1, tmp)
}

/**
归并排序 -> 分治法
*/
func sortt(arr []int, left, right int, tmp []int) {
	if left < right {
		mid := (left + right) / 2
		sortt(arr, left, mid, tmp)        // 分治法 -> 先分
		sortt(arr, mid+1, right, tmp)     // 分治法 -> 先分
		merge(arr, left, mid, right, tmp) // 分治法 -> 后治
	}
}

func merge(arr []int, left, mid, right int, tmp []int) {
	i, j := left, mid+1
	t := 0
	for i <= mid && j <= right { // 合并两个有序数组
		if arr[i] <= arr[j] {
			tmp[t] = arr[i]
			i++
		} else {
			tmp[t] = arr[j]
			j++
		}
		t++
	}
	fmt.Println(tmp, i, mid)
	for i <= mid { // 将左边剩余的数加入数组中
		tmp[t] = arr[i]
		t++
		i++
	}

	for j <= right { // 将右边剩余的数加入数组中
		tmp[t] = arr[j]
		t++
		j++
	}
	t = 0
	for left <= right { // 将临时数组复制回原数组
		arr[left] = tmp[t]
		left++
		t++
	}
	fmt.Println(arr)
}

// 魔改版分治法
func reversePairs(arr []int) int {
	l := len(arr)
	tmp := make([]int, l)
	return sorttCount(arr, 0, l-1, tmp)
}

func sorttCount(arr []int, left, right int, tmp []int) int {
	if left >= right {
		return 0
	}
	mid := (left + right) / 2
	return sorttCount(arr, left, mid, tmp) + // 分治法 -> 先分
		sorttCount(arr, mid+1, right, tmp) + // 分治法 -> 先分
		mergeCount(arr, left, mid, right, tmp) // 分治法 -> 后治
}

// 魔改版分治法
func mergeCount(arr []int, left, mid, right int, tmp []int) int {
	i, j := left, mid+1
	t := 0
	res := 0
	for i <= mid && j <= right { // 合并两个有序数组
		if arr[i] <= arr[j] {
			tmp[t] = arr[i]
			i++
			res += j - (mid + 1) // 因为是已经排序好的数组, 所以当发现小于时, 应该比前面所有都小
			// 此时 arr[i] 比 arr[j] 小，把 arr[i] 对应的数加入tmp中，
			// 并考虑它对逆序对总数的贡献为 j 相对 右边 首位置的偏移 mid+1（即右边只有x个数比 arr[j] 小，
			// 所以只有它和x个数构成逆序数列），以此类推。
		} else {
			tmp[t] = arr[j] // 大于时不需要管
			j++
		}
		t++
	}
	for i <= mid {
		tmp[t] = arr[i]
		t++
		i++
		res += j - (mid + 1) // 此处也相同
	}

	for j <= right {
		tmp[t] = arr[j]
		t++
		j++
	}
	t = 0
	for left <= right {
		arr[left] = tmp[t]
		left++
		t++
	}
	return res
}

func main() {
	fmt.Println(reversePairs2([]int{7, 5, 6, 4}))
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))

	fmt.Println("\nTest Merge Sort")
	arr := []int{2, 3, 4, 1, 5, 6}
	sort(arr)
	fmt.Println(arr)
	arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	sort(arr)
	fmt.Println(arr)
}
