package main

import "fmt"

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	temp := make([]int, k)
	copy(temp, nums[l-k:])
	for i := l - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	copy(nums, temp)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	// rotate(nums, 3)
	rotateByReverse(nums, 3)
	fmt.Println(nums)
}

func rotateByReverse(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}
