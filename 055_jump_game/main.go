package main

import "fmt"

func canJump(nums []int) bool {
	dp := 0
	for i := range nums {
		if i > dp {
			return false
		}
		dp = max(dp, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	data := [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
	}
	for i := range data {
		fmt.Println(canJump(data[i]))
	}
}
