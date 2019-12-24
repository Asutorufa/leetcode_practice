package main_test

import "testing"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
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
func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
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

func BenchmarkMedia(b *testing.B) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindMedianSortedArrays2(nums1, nums2)
	}
}

func BenchmarkMedia2(b *testing.B) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindMedianSortedArrays(nums1, nums2)
	}
}
