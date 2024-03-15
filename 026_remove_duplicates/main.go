package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow, fast := 1, 1
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func main() {
	t := []int{1, 1, 2, 3}
	// copy(t[0:], t[1:])
	l := removeDuplicates(t)
	fmt.Println(l, ":", t)
}
