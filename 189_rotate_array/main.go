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
	rotate(nums, 3)
	fmt.Println(nums)
}
