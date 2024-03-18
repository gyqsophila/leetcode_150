package main

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	count := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[count-2] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func removeDuplicatesWithTwoPoint(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	slow, fast := 2, 2
	for fast < len(nums) {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
