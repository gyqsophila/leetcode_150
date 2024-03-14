package main

import "fmt"

func removeElement(nums []int, val int) int {
	return twoPointFromEndToStart(nums, val)
}

func twoPointFromEndToStart(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}

func main() {
	nums := []int{3, 2, 2, 3}
	l := removeElement(nums, 3)
	fmt.Println(nums)
	fmt.Println(l)
}

func twoPointFromStartToEnd(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
