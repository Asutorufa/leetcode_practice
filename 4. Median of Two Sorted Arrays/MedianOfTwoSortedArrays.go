package main

import "log"

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0

Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	x, y := 0, 0
	nums := []int{}
	for {
		if x < len(nums1) && y < len(nums2) {
			if nums1[x] < nums2[y] {
				nums = append(nums, nums1[x])
				x++
				continue
			} else {
				nums = append(nums, nums2[y])
				y++
				continue
			}
		} else if x >= len(nums1) && y <= len(nums2) {
			nums = append(nums, nums2[y:]...)
			break
		} else if x <= len(nums1) && y >= len(nums2) {
			nums = append(nums, nums1[x:]...)
			break
		} else {
			break
		}
	}
	if len(nums)%2 == 1 {
		return float64(nums[len(nums)/2])
	}
	return (float64(nums[len(nums)/2-1]) + float64(nums[len(nums)/2])) / 2
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	log.Println(findMedianSortedArrays(nums1, nums2))
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	log.Println(findMedianSortedArrays(nums1, nums2))
}
