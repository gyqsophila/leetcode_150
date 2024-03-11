package main

import "testing"

type Data struct {
	Nums1       []int
	Nums1Length int
	Nums2       []int
	Nums2Length int
	Want        []int
}

var (
	data = []Data{
		{
			Nums1:       []int{1, 2, 3, 0, 0, 0},
			Nums1Length: 3,
			Nums2:       []int{2, 5, 6},
			Nums2Length: 3,
			Want:        []int{1, 2, 2, 3, 5, 6},
		},
		{
			Nums1:       []int{1},
			Nums1Length: 1,
			Nums2:       []int{},
			Nums2Length: 0,
			Want:        []int{1},
		},
		{
			Nums1:       []int{0},
			Nums1Length: 0,
			Nums2:       []int{1},
			Nums2Length: 1,
			Want:        []int{1},
		},
	}
)

func TestMergeArray(t *testing.T) {
	for _, example := range data {
		merge(example.Nums1, example.Nums1Length, example.Nums2, example.Nums2Length)
		for i := range example.Nums1 {
			if example.Nums1[i] != example.Want[i] {
				t.Errorf("nums[%d]!=want[%d]", i, i)
				t.Logf("nums1: %v", example.Nums1)
				t.Logf("want: %v", example.Want)
			}
		}
	}
}
