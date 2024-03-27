package main

import "fmt"

var (
	nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roms = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

func intToRoman(num int) string {

	var str string
	i := 0
	for num > 0 {
		if num/nums[i] > 0 {
			str += roms[i]
			num -= nums[i]
		} else {
			i++
		}
	}
	return str
}

func main() {
	ints := []int{3, 9, 58, 1994}
	for i := range ints {
		fmt.Println(intToRoman(ints[i]))
	}
}
