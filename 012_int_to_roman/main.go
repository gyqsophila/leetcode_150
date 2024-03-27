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

var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func mod(num int) string {

	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]

}

func main() {
	ints := []int{3, 9, 58, 1994}
	for i := range ints {
		fmt.Println(mod(ints[i]))
	}
}
