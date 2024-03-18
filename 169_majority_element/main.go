package main

import "sort"

func main() {

}

func majorityElementWithEnum(nums []int) int {
	statistic := map[int]int{}
	num, max := nums[0], 1
	for _, n := range nums {
		statistic[n]++
		if statistic[n] > max {
			num, max = n, statistic[n]
		}
	}
	return num
}

func majorityElementWithSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement(nums []int) int {
	much, time := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == much {
			time++
		} else {
			if time == 0 {
				much = nums[i]
				time++
			} else {
				time--
			}
		}
	}
	return much
}
