package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	// mergeByLibrary(nums1, m, nums2, n)
	mergeFromEndToStart(nums1, m, nums2, n)
}

// call go official library func to calcute
func mergeByLibrary(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}
	sort.Ints(nums1)
}

// from end to strt of nums1
// and also pick out bigger one digit from nums1 and nums2
// if once picked ,then take one front of related pointer
// and also take one front position of put in
// and finally, if nums1 get to 0 position first, then copy left num2 to beginnning of nums1
func mergeFromEndToStart(nums1 []int, m int, nums2 []int, n int) {
	l1 := m - 1
	l2 := n - 1
	len := m + n - 1
	for l1 >= 0 && l2 >= 0 {
		if nums1[l1] > nums2[l2] {
			nums1[len] = nums1[l1]
			l1--
		} else {
			nums1[len] = nums2[l2]
			l2--
		}
		len--
	}
	fmt.Println("before copy: ", nums1)
	copy(nums1, nums2[0:l2+1])
}
