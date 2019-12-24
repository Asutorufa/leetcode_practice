package main

import (
	"log"
)

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

// recommend for this video: https://www.youtube.com/watch?time_continue=1&v=LPFhl65R7ww, that's clear.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	if n == 0 {
		panic("ValueError")
	}
	imin, imax, halfLen := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i
		if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			var maxOfLeft int
			var minOfRight int
			if i == 0 {
				maxOfLeft = nums2[j-1]
			} else if j == 0 {
				maxOfLeft = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					maxOfLeft = nums1[i-1]
				} else {
					maxOfLeft = nums2[j-1]
				}
			}
			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}
			if i == m {
				minOfRight = nums2[j]
			} else if j == n {
				minOfRight = nums1[i]
			} else {
				if nums1[i] < nums2[j] {
					minOfRight = nums1[i]
				} else {
					minOfRight = nums2[j]
				}
			}
			return (float64(maxOfLeft) + float64(minOfRight)) / 2.0
		}
	}
	return 0
}

/*

	wrong answer,the time complexity is O(m+n)

*/
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
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
	log.Println(findMedianSortedArrays2(nums1, nums2))
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	log.Println(findMedianSortedArrays(nums1, nums2))
	log.Println(findMedianSortedArrays2(nums1, nums2))
	nums1 = []int{1, 2}
	nums2 = []int{-1, 3}
	log.Println(findMedianSortedArrays(nums1, nums2))
	log.Println(findMedianSortedArrays2(nums1, nums2))
	nums1 = []int{1, 3, 5, 7, 9, 11}
	nums2 = []int{-5, -3, -1, 2, 4, 6}
	log.Println(findMedianSortedArrays(nums1, nums2))
	log.Println(findMedianSortedArrays2(nums1, nums2))
}
