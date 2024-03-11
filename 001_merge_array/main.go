package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	mergeByLibrary(nums1, m, nums2, n)
}

// call go official library func to calcute
func mergeByLibrary(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}
	sort.Ints(nums1)
}
